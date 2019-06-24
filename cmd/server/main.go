package main

import (
	"context"
	"database/sql"
	"fmt"
	"contacts1/pkg/pb"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/server"

	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	"github.com/infobloxopen/atlas-app-toolkit/health"
)

func main() {
	doneC := make(chan error)
	logger := NewLogger()

	if viper.GetBool("internal.enable") {
		go func() { doneC <- ServeInternal(logger) }()
	}

	go func() { doneC <- ServeExternal(logger) }()

	if err := <-doneC; err != nil {
		logger.Fatal(err)
	}
}

func NewLogger() *logrus.Logger {
	logger := logrus.StandardLogger()

	// Set the log level on the default logger based on command line flag
	logLevels := map[string]logrus.Level{
		"debug":   logrus.DebugLevel,
		"info":    logrus.InfoLevel,
		"warning": logrus.WarnLevel,
		"error":   logrus.ErrorLevel,
		"fatal":   logrus.FatalLevel,
		"panic":   logrus.PanicLevel,
	}
	if level, ok := logLevels[viper.GetString("logging.level")]; !ok {
		logger.Errorf("Invalid %q provided for log level", viper.GetString("logging.level"))
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}

	return logger
}

// ServeInternal builds and runs the server that listens on InternalAddress
func ServeInternal(logger *logrus.Logger) error {
	healthChecker := health.NewChecksHandler(
		viper.GetString("internal.health"),
		viper.GetString("internal.readiness"),
	)
	healthChecker.AddReadiness("DB ready check", dbReady)
	healthChecker.AddLiveness("ping", health.HTTPGetCheck(
		fmt.Sprint("http://", viper.GetString("internal.address"), ":", viper.GetString("internal.port"), "/ping"), time.Minute),
	)

	s, err := server.NewServer(
		// register our health checks
		server.WithHealthChecks(healthChecker),
		// this endpoint will be used for our health checks
		server.WithHandler("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("pong"))
		})),
	)
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))
	if err != nil {
		return err
	}

	logger.Debugf("serving internal http at %q", fmt.Sprintf("%s:%s", viper.GetString("internal.address"), viper.GetString("internal.port")))
	return s.Serve(nil, l)
}

// ServeExternal builds and runs the server that listens on ServerAddress and GatewayAddress
func ServeExternal(logger *logrus.Logger) error {

	if viper.GetString("database.dsn") == "" {
		setDBConnection()
	}
	grpcServer, err := NewGRPCServer(logger, viper.GetString("database.dsn"))
	if err != nil {
		logger.Fatalln(err)
	}

	s, err := server.NewServer(
		server.WithGrpcServer(grpcServer),
		server.WithGateway(
			gateway.WithGatewayOptions(
				runtime.WithForwardResponseOption(forwardResponseOption),
			),
			gateway.WithServerAddress(fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port"))),
			gateway.WithEndpointRegistration(viper.GetString("gateway.endpoint"), pb.RegisterContacts1HandlerFromEndpoint),
		),
		server.WithHandler("/swagger/", NewSwaggerHandler(viper.GetString("gateway.swaggerFile"))),
	)
	if err != nil {
		logger.Fatalln(err)
	}

	grpcL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("server.address"), viper.GetString("server.port")))
	if err != nil {
		logger.Fatalln(err)
	}

	httpL, err := net.Listen("tcp", fmt.Sprintf("%s:%s", viper.GetString("gateway.address"), viper.GetString("gateway.port")))
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Printf("serving gRPC at %s:%s", viper.GetString("server.address"), viper.GetString("server.port"))
	logger.Printf("serving http at %s:%s", viper.GetString("gateway.address"), viper.GetString("gateway.port"))

	return s.Serve(grpcL, httpL)
}

func init() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(viper.GetString("config.source"))
	if viper.GetString("config.file") != "" {
		log.Printf("Serving from configuration file: %s", viper.GetString("config.file"))
		viper.SetConfigName(viper.GetString("config.file"))
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("cannot load configuration: %v", err)
		}
	} else {
		log.Printf("Serving from default values, environment variables, and/or flags")
	}
	resource.RegisterApplication(viper.GetString("app.id"))
	resource.SetPlural()
}

func forwardResponseOption(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	return nil
}

func dbReady() error {
	if viper.GetString("database.dsn") == "" {
		setDBConnection()
	}
	db, err := sql.Open(viper.GetString("database.type"), viper.GetString("database.dsn"))
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Ping()
}

// setDBConnection sets the db connection string
func setDBConnection() {
	viper.Set("database.dsn", fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		viper.GetString("database.address"), viper.GetString("database.port"),
		viper.GetString("database.user"), viper.GetString("database.password"),
		viper.GetString("database.ssl"), viper.GetString("database.name")))
}
