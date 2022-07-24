package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		check func(error) bool
	}{
		{
			name:  "Cancelled",
			err:   Canceled("reason_Cancelled", "message_Cancelled"),
			check: IsCanceled,
		},
		{
			name:  "Unknown",
			err:   Unknown("reason_Unknown", "message_Unknown"),
			check: IsUnknown,
		},
		{
			name:  "InvalidArgument",
			err:   InvalidArgument("reason_InvalidArgument", "message_InvalidArgument"),
			check: IsInvalidArgument,
		},
		{
			name:  "DeadlineExceeded",
			err:   DeadlineExceeded("reason_DeadlineExceeded", "message_DeadlineExceeded"),
			check: IsDeadlineExceeded,
		},
		{
			name:  "NotFound",
			err:   NotFound("reason_NotFound", "message_NotFound"),
			check: IsNotFound,
		},
		{
			name:  "AlreadyExists",
			err:   AlreadyExists("reason_AlreadyExists", "message_AlreadyExists"),
			check: IsAlreadyExists,
		},
		{
			name:  "PermissionDenied",
			err:   PermissionDenied("reason_PermissionDenied", "message_PermissionDenied"),
			check: IsPermissionDenied,
		},
		{
			name:  "ResourceExhausted",
			err:   ResourceExhausted("reason_ResourceExhausted", "message_ResourceExhausted"),
			check: IsResourceExhausted,
		},
		{
			name:  "FailedPrecondition",
			err:   FailedPrecondition("reason_FailedPrecondition", "message_FailedPrecondition"),
			check: IsFailedPrecondition,
		},
		{
			name:  "Aborted",
			err:   Aborted("reason_Aborted", "message_Aborted"),
			check: IsAborted,
		},
		{
			name:  "OutOfRange",
			err:   OutOfRange("reason_OutOfRange", "message_OutOfRange"),
			check: IsOutOfRange,
		},
		{
			name:  "Unimplemented",
			err:   Unimplemented("reason_Unimplemented", "message_Unimplemented"),
			check: IsUnimplemented,
		},
		{
			name:  "Internal",
			err:   Internal("reason_Internal", "message_Internal"),
			check: IsInternal,
		},
		{
			name:  "Unavailable",
			err:   Unavailable("reason_Unavailable", "message_Unavailable"),
			check: IsUnavailable,
		},
		{
			name:  "DataLoss",
			err:   DataLoss("reason_DataLoss", "message_DataLoss"),
			check: IsDataLoss,
		},
		{
			name:  "Unauthenticated",
			err:   Unauthenticated("reason_Unauthenticated", "message_Unauthenticated"),
			check: IsUnauthenticated,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, test.check(test.err))
		})
	}
}
