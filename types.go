package errors

import (
	"google.golang.org/grpc/codes"
)

// Canceled creates an Error with codes.Canceled which mapping to HTTP code 499.
func Canceled(reason, message string) *Error {
	return New(codes.Canceled, reason, message)
}

// IsCanceled checks if an error is an Error with codes.Canceled which mapping to HTTP code 499.
// This function supports wrapped errors.
func IsCanceled(err error) bool {
	return Code(err) == codes.Canceled
}

// Unknown creates an Error with codes.Unknown which mapping to HTTP code 500.
func Unknown(reason, message string) *Error {
	return New(codes.Unknown, reason, message)
}

// IsUnknown checks if an error is an Error with codes.Unknown which mapping to HTTP code 500.
// It supports wrapped errors.
func IsUnknown(err error) bool {
	return Code(err) == codes.Unknown
}

// InvalidArgument creates an Error with codes.InvalidArgument which mapping to HTTP code 400.
func InvalidArgument(reason, message string) *Error {
	return New(codes.InvalidArgument, reason, message)
}

// IsInvalidArgument checks if an error is an Error with codes.InvalidArgument which mapping to HTTP code 400.
// It supports wrapped errors.
func IsInvalidArgument(err error) bool {
	return Code(err) == codes.InvalidArgument
}

// DeadlineExceeded creates an Error with codes.DeadlineExceeded which mapping to HTTP code 504.
func DeadlineExceeded(reason, message string) *Error {
	return New(codes.DeadlineExceeded, reason, message)
}

// IsDeadlineExceeded checks if an error is an Error with codes.DeadlineExceeded which mapping to HTTP code 504.
// It supports wrapped errors.
func IsDeadlineExceeded(err error) bool {
	return Code(err) == codes.DeadlineExceeded
}

// NotFound creates an Error with codes.NotFound which mapping to HTTP code 404.
func NotFound(reason, message string) *Error {
	return New(codes.NotFound, reason, message)
}

// IsNotFound checks if an error is an Error with codes.NotFound which mapping to HTTP code 404.
// It supports wrapped errors.
func IsNotFound(err error) bool {
	return Code(err) == codes.NotFound
}

// AlreadyExists creates an Error with codes.AlreadyExists which mapping to HTTP code 409.
func AlreadyExists(reason, message string) *Error {
	return New(codes.AlreadyExists, reason, message)
}

// IsAlreadyExists checks if an error is an Error with codes.AlreadyExists which mapping to HTTP code 409.
// It supports wrapped errors.
func IsAlreadyExists(err error) bool {
	return Code(err) == codes.AlreadyExists
}

// PermissionDenied creates an Error with codes.PermissionDenied which mapping to HTTP code 403.
func PermissionDenied(reason, message string) *Error {
	return New(codes.PermissionDenied, reason, message)
}

// IsPermissionDenied checks if an error is an Error with codes.PermissionDenied which mapping to HTTP code 403.
// It supports wrapped errors.
func IsPermissionDenied(err error) bool {
	return Code(err) == codes.PermissionDenied
}

// ResourceExhausted creates an Error with codes.ResourceExhausted which mapping to HTTP code 429.
func ResourceExhausted(reason, message string) *Error {
	return New(codes.ResourceExhausted, reason, message)
}

// IsResourceExhausted checks if an error is an Error with codes.ResourceExhausted which mapping to HTTP code 429.
// It supports wrapped errors.
func IsResourceExhausted(err error) bool {
	return Code(err) == codes.ResourceExhausted
}

// FailedPrecondition creates an Error with codes.FailedPrecondition which mapping to HTTP code 400.
func FailedPrecondition(reason, message string) *Error {
	return New(codes.FailedPrecondition, reason, message)
}

// IsFailedPrecondition checks if an error is an Error with codes.FailedPrecondition which mapping to HTTP code 400.
// It supports wrapped errors.
func IsFailedPrecondition(err error) bool {
	return Code(err) == codes.FailedPrecondition
}

// Aborted creates an Error with codes.Aborted which mapping to HTTP code 409.
func Aborted(reason, message string) *Error {
	return New(codes.Aborted, reason, message)
}

// IsAborted checks if an error is an Error with codes.Aborted which mapping to HTTP code 409.
// It supports wrapped errors.
func IsAborted(err error) bool {
	return Code(err) == codes.Aborted
}

// OutOfRange creates an Error with codes.OutOfRange which mapping to HTTP code 400.
func OutOfRange(reason, message string) *Error {
	return New(codes.OutOfRange, reason, message)
}

// IsOutOfRange checks if an error is an Error with codes.OutOfRange which mapping to HTTP code 400.
// It supports wrapped errors.
func IsOutOfRange(err error) bool {
	return Code(err) == codes.OutOfRange
}

// Unimplemented creates an Error with codes.Unimplemented which mapping to HTTP code 501.
func Unimplemented(reason, message string) *Error {
	return New(codes.Unimplemented, reason, message)
}

// IsUnimplemented checks if an error is an Error with codes.Unimplemented which mapping to HTTP code 501.
// It supports wrapped errors.
func IsUnimplemented(err error) bool {
	return Code(err) == codes.Unimplemented
}

// Internal creates an Error with codes.Internal which mapping to HTTP code 500.
func Internal(reason, message string) *Error {
	return New(codes.Internal, reason, message)
}

// IsInternal checks if an error is an Error with codes.Internal which mapping to HTTP code 500.
// It supports wrapped errors.
func IsInternal(err error) bool {
	return Code(err) == codes.Internal
}

// Unavailable creates an Error with codes.Unavailable which mapping to HTTP code 503.
func Unavailable(reason, message string) *Error {
	return New(codes.Unavailable, reason, message)
}

// IsUnavailable checks if an error is an Error with codes.Unavailable which mapping to HTTP code 503.
// It supports wrapped errors.
func IsUnavailable(err error) bool {
	return Code(err) == codes.Unavailable
}

// DataLoss creates an Error with codes.DataLoss which mapping to HTTP code 500.
func DataLoss(reason, message string) *Error {
	return New(codes.DataLoss, reason, message)
}

// IsDataLoss checks if an error is an Error with codes.DataLoss which mapping to HTTP code 500.
// It supports wrapped errors.
func IsDataLoss(err error) bool {
	return Code(err) == codes.DataLoss
}

// Unauthenticated creates an Error with codes.Unauthenticated which mapping to HTTP code 401.
func Unauthenticated(reason, message string) *Error {
	return New(codes.Unauthenticated, reason, message)
}

// IsUnauthenticated checks if an error is an Error with codes.Unauthenticated which mapping to HTTP code 401.
// It supports wrapped errors.
func IsUnauthenticated(err error) bool {
	return Code(err) == codes.Unauthenticated
}
