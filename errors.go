package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	gencode "google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//goland:noinspection GoUnusedConst
const (
	// UnknownReason is unknown reason for error info.
	UnknownReason = "UNKNOWN"
	// SupportPackageIsVersion1 is used for generation validating,
	// and this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
	// ClientClosedHttpCode is non-standard http status code,
	// which defined by nginx.
	// https://httpstatus.in/499/
	ClientClosedHttpCode = 499
)

// Error is an error with Status and an optional cause.
type Error struct {
	Status
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"error: code = %s reason = %q message = %q metadata = %q cause = %v",
		e.Code, e.Reason, e.Message, e.Metadata, e.cause,
	)
}

func (e *Error) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "error: code = %s reason = %q message = %q metadata = %q cause = %+v",
				e.Code, e.Reason, e.Message, e.Metadata, e.cause,
			)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Unwrap() error { return e.cause }

func (e *Error) MarshalJSON() ([]byte, error) {
	reason, _ := json.Marshal(e.Reason)
	message, _ := json.Marshal(e.Message)
	if len(e.Metadata) == 0 {
		return []byte(fmt.Sprintf(
			`{"code":"%s","reason":%s,"message":%s}`,
			e.Code.String(), string(reason), string(message),
		)), nil
	}
	metadata, _ := json.Marshal(e.Metadata)
	return []byte(fmt.Sprintf(
		`{"code":"%s","reason":%s,"message":%s,"metadata":%s}`,
		e.Code.String(), string(reason), string(message), metadata,
	)), nil
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

// WithCause returns a new Error with the underlying cause.
func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

// WithMetadata returns a new Error with a new key-value metadata map.
// The metadata of origin Error would be dropped.
func (e *Error) WithMetadata(md map[string]string) *Error {
	err := Clone(e)
	err.Metadata = md
	return err
}

// WithForceHttpCode returns a new Error that has set an HTTP status code the Error would be mapped to.
func (e *Error) WithForceHttpCode(code int32) *Error {
	err := Clone(e)
	err.ForceHttpCode = code
	return err
}

// GRPCStatus returns the Status represented by the Error.
func (e *Error) GRPCStatus() *status.Status {
	s, _ := status.New(codes.Code(e.Code), e.Message).
		WithDetails(&errdetails.ErrorInfo{
			Reason:   e.Reason,
			Metadata: e.Metadata,
		})
	return s
}

func (e *Error) HttpCode() int32 {
	if e.ForceHttpCode != 0 {
		return e.ForceHttpCode
	}
	switch codes.Code(e.Code) {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return ClientClosedHttpCode
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}

// AppendMetadata returns a new Error, with a new metadata key-value pair appended or replaced.
func (e *Error) AppendMetadata(key, value string) *Error {
	err := Clone(e)
	err.Metadata[key] = value
	return err
}

// New returns an Error for the code, reason and message.
func New(code codes.Code, reason, message string) *Error {
	return &Error{
		Status: Status{
			Code:    gencode.Code(code),
			Message: message,
			Reason:  reason,
		},
	}
}

// Newf is same as New(code, reason, fmt.Sprintf(format, a...))
func Newf(code codes.Code, reason, format string, a ...any) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code codes.Code, reason, format string, a ...any) error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Code returns the code for an error.
// It supports wrapped errors.
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	return codes.Code(FromError(err).Code)
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}

// Clone deep clones Error to a new Error.
func Clone(err *Error) *Error {
	metadata := make(map[string]string, len(err.Metadata))
	for k, v := range err.Metadata {
		metadata[k] = v
	}
	return &Error{
		cause: err.cause,
		Status: Status{
			Code:     err.Code,
			Reason:   err.Reason,
			Message:  err.Message,
			Metadata: metadata,
		},
	}
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if ok {
		ret := New(gs.Code(), UnknownReason, gs.Message())
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				ret.Reason = d.Reason
				return ret.WithMetadata(d.Metadata)
			}
		}
		return ret
	}
	return New(codes.Unknown, UnknownReason, err.Error())
}
