# build the server binary
FROM golang:1.10.0 AS builder
LABEL stage=server-intermediate
WORKDIR /go/src/github.com/gshirazi/contacts1
COPY . .
#RUN export GOPATH=$GOPATH:/Users/gshirazi/go
RUN echo "GOPATH is $GOPATH"
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server

# copy the server binary from builder stage; run the server binary
FROM alpine:latest AS runner
#FROM postgres:latest as runner
WORKDIR /bin
COPY --from=builder /go/src/github.com/gshirazi/contacts1/bin/server .
COPY pkg/pb/*.swagger.json tmp/www/swagger
COPY --from=builder /go/src/github.com/gshirazi/contacts1/db/migrations /db/migrations/

# RUN su - postgres -c "cat /var/lib/postgresql/log.log"
# RUN su - postgres -c "find / 2>/dev/null | grep postg"
# RUN su - postgres -c  "cat /var/lib/postgresql/data/postgresql.conf | grep tcp"
# RUN su - postgres -c "pg_ctl status -D /var/lib/postgresql/data -l /var/lib/postgresql/log.log"
#ENTRYPOINT ["/bin/sh echo hello"]
ENTRYPOINT ["server", "--gateway.swaggerFile", "www/contacts1.swagger.json"]