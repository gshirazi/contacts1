// Code generated by protoc-gen-validate
// source: contacts1/pkg/pb/service.proto
// DO NOT EDIT!!!

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on VersionResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *VersionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	return nil
}

// VersionResponseValidationError is the validation error returned by
// VersionResponse.Validate if the designated constraints aren't met.
type VersionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VersionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VersionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VersionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VersionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VersionResponseValidationError) ErrorName() string { return "VersionResponseValidationError" }

// Error satisfies the builtin error interface
func (e VersionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVersionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VersionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VersionResponseValidationError{}

// Validate checks the field values on Contact with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Contact) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Email

	// no validation rules for HomeAddress

	return nil
}

// ContactValidationError is the validation error returned by Contact.Validate
// if the designated constraints aren't met.
type ContactValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContactValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContactValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContactValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContactValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContactValidationError) ErrorName() string { return "ContactValidationError" }

// Error satisfies the builtin error interface
func (e ContactValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContact.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContactValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContactValidationError{}

// Validate checks the field values on ReverseRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReverseRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReverseRequestValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReverseRequestValidationError is the validation error returned by
// ReverseRequest.Validate if the designated constraints aren't met.
type ReverseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReverseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReverseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReverseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReverseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReverseRequestValidationError) ErrorName() string { return "ReverseRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReverseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReverseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReverseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReverseRequestValidationError{}

// Validate checks the field values on ReverseResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ReverseResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// ReverseResponseValidationError is the validation error returned by
// ReverseResponse.Validate if the designated constraints aren't met.
type ReverseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReverseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReverseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReverseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReverseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReverseResponseValidationError) ErrorName() string { return "ReverseResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReverseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReverseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReverseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReverseResponseValidationError{}

// Validate checks the field values on CreateContactRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateContactRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateContactRequestValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateContactRequestValidationError is the validation error returned by
// CreateContactRequest.Validate if the designated constraints aren't met.
type CreateContactRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateContactRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateContactRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateContactRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateContactRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateContactRequestValidationError) ErrorName() string {
	return "CreateContactRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateContactRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateContactRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateContactRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateContactRequestValidationError{}

// Validate checks the field values on CreateContactResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateContactResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateContactResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateContactResponseValidationError is the validation error returned by
// CreateContactResponse.Validate if the designated constraints aren't met.
type CreateContactResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateContactResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateContactResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateContactResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateContactResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateContactResponseValidationError) ErrorName() string {
	return "CreateContactResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateContactResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateContactResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateContactResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateContactResponseValidationError{}

// Validate checks the field values on ReadContactRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadContactRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// ReadContactRequestValidationError is the validation error returned by
// ReadContactRequest.Validate if the designated constraints aren't met.
type ReadContactRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadContactRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadContactRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadContactRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadContactRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadContactRequestValidationError) ErrorName() string {
	return "ReadContactRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadContactRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadContactRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadContactRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadContactRequestValidationError{}

// Validate checks the field values on ReadContactResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadContactResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadContactResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadContactResponseValidationError is the validation error returned by
// ReadContactResponse.Validate if the designated constraints aren't met.
type ReadContactResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadContactResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadContactResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadContactResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadContactResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadContactResponseValidationError) ErrorName() string {
	return "ReadContactResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadContactResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadContactResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadContactResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadContactResponseValidationError{}

// Validate checks the field values on UpdateContactRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateContactRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateContactRequestValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateContactRequestValidationError is the validation error returned by
// UpdateContactRequest.Validate if the designated constraints aren't met.
type UpdateContactRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateContactRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateContactRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateContactRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateContactRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateContactRequestValidationError) ErrorName() string {
	return "UpdateContactRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateContactRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateContactRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateContactRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateContactRequestValidationError{}

// Validate checks the field values on UpdateContactResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateContactResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateContactResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateContactResponseValidationError is the validation error returned by
// UpdateContactResponse.Validate if the designated constraints aren't met.
type UpdateContactResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateContactResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateContactResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateContactResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateContactResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateContactResponseValidationError) ErrorName() string {
	return "UpdateContactResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateContactResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateContactResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateContactResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateContactResponseValidationError{}

// Validate checks the field values on DeleteContactRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteContactRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteContactRequestValidationError is the validation error returned by
// DeleteContactRequest.Validate if the designated constraints aren't met.
type DeleteContactRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteContactRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteContactRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteContactRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteContactRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteContactRequestValidationError) ErrorName() string {
	return "DeleteContactRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteContactRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteContactRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteContactRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteContactRequestValidationError{}

// Validate checks the field values on DeleteContactResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteContactResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteContactResponseValidationError is the validation error returned by
// DeleteContactResponse.Validate if the designated constraints aren't met.
type DeleteContactResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteContactResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteContactResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteContactResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteContactResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteContactResponseValidationError) ErrorName() string {
	return "DeleteContactResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteContactResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteContactResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteContactResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteContactResponseValidationError{}