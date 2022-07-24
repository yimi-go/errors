package errors

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/grpc_testing"
)

type testError struct{ message string }

func (e *testError) Error() string { return e.message }

func TestError(t *testing.T) {
	var base *Error
	err := Newf(codes.InvalidArgument, "reason", "message")
	err2 := Newf(codes.InvalidArgument, "reason", "message")
	err3 := err.WithMetadata(map[string]string{
		"foo": "bar",
	})
	err4 := err3.AppendMetadata("hello", "world")
	werr := fmt.Errorf("wrap %w", err)

	assert.NotErrorIs(t, new(Error), err)
	assert.ErrorIs(t, err, werr)
	assert.Error(t, err2, werr)
	assert.True(t, errors.As(err, &base))
	//assert.True(t, IsInvalidArgument(err))
	assert.Equal(t, err3.Reason, Reason(err))
	assert.Equal(t, "bar", err3.Metadata["foo"])
	assert.Equal(t, "bar", err4.Metadata["foo"])
	assert.Equal(t, "world", err4.Metadata["hello"])
	gs := err.GRPCStatus()
	se := FromError(gs.Err())
	assert.Equal(t, "reason", se.Reason)
	gs2, _ := status.New(codes.InvalidArgument, "bad request").WithDetails(&grpc_testing.Empty{})
	se2 := FromError(gs2.Err())
	assert.Equal(t, codes.InvalidArgument, codes.Code(se2.Code))
	assert.Nil(t, FromError(nil))
	e := FromError(errors.New("text"))
	assert.Equal(t, codes.Unknown, codes.Code(e.Code))
}

func TestClone(t *testing.T) {
	e := &Error{}
	e2 := Clone(e)
	assert.NotNil(t, e2)
	assert.NotNil(t, e2.Message)
}

func TestError_Is(t *testing.T) {
	tests := []struct {
		name string
		e    *Error
		err  error
		want bool
	}{
		{
			name: "true",
			e:    &Error{Status: Status{Code: code.Code_NOT_FOUND, Reason: "test"}},
			err:  New(codes.NotFound, "test", ""),
			want: true,
		},
		{
			name: "false",
			e:    &Error{Status: Status{Reason: "test"}},
			err:  errors.New("test"),
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok := test.e.Is(test.err)
			assert.Equal(t, test.want, ok)
		})
	}
}

func TestOther(t *testing.T) {
	assert.Equal(t, codes.OK, Code(nil), "Code(nil)")
	assert.Equal(t, codes.Unknown, Code(errors.New("test")), `Code(errors.New("test"))`)
	err := Errorf(codes.Code(10001), "test code 10001", "message")
	assert.Equal(t, codes.Code(10001), Code(err))
	assert.Equal(t, "test code 10001", Reason(err))
	err2 := New(codes.Unknown, "", "").AppendMetadata("foo", "bar")
	assert.Equal(t, "bar", err2.Metadata["foo"])
	assert.Equal(t, UnknownReason, Reason(nil))
}

func TestCause(t *testing.T) {
	testErr := &testError{message: "test"}
	err := New(codes.InvalidArgument, "foo", "bar").WithCause(testErr)
	assert.ErrorIs(t, err, testErr)
	te := &testError{}
	assert.True(t, errors.As(err, &te))
	assert.Equal(t, testErr.message, te.message)
}

func TestError_WithHttpCode(t *testing.T) {
	err := New(codes.Unknown, "", "").WithForceHttpCode(http.StatusBadRequest)
	assert.Equal(t, http.StatusBadRequest, int(err.HttpCode()))
}

func TestError_HttpCode(t *testing.T) {
	tests := []struct {
		force int32
		code  code.Code
		want  int32
	}{
		{
			force: http.StatusNotModified,
			code:  code.Code_OK,
			want:  http.StatusNotModified,
		},
		{code: code.Code_OK, want: http.StatusOK},
		{code: code.Code_CANCELLED, want: ClientClosedHttpCode},
		{code: code.Code_UNKNOWN, want: http.StatusInternalServerError},
		{code: code.Code_INVALID_ARGUMENT, want: http.StatusBadRequest},
		{code: code.Code_DEADLINE_EXCEEDED, want: http.StatusGatewayTimeout},
		{code: code.Code_NOT_FOUND, want: http.StatusNotFound},
		{code: code.Code_ALREADY_EXISTS, want: http.StatusConflict},
		{code: code.Code_PERMISSION_DENIED, want: http.StatusForbidden},
		{code: code.Code_UNAUTHENTICATED, want: http.StatusUnauthorized},
		{code: code.Code_RESOURCE_EXHAUSTED, want: http.StatusTooManyRequests},
		{code: code.Code_FAILED_PRECONDITION, want: http.StatusBadRequest},
		{code: code.Code_ABORTED, want: http.StatusConflict},
		{code: code.Code_OUT_OF_RANGE, want: http.StatusBadRequest},
		{code: code.Code_UNIMPLEMENTED, want: http.StatusNotImplemented},
		{code: code.Code_INTERNAL, want: http.StatusInternalServerError},
		{code: code.Code_UNAVAILABLE, want: http.StatusServiceUnavailable},
		{code: code.Code_DATA_LOSS, want: http.StatusInternalServerError},
		{code: code.Code(599), want: http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_%v", tt.force, tt.code), func(t *testing.T) {
			e := &Error{
				Status: Status{
					Code:          tt.code,
					ForceHttpCode: tt.force,
				},
			}
			assert.Equal(t, tt.want, e.HttpCode())
		})
	}
}

func TestError_MarshalJSON(t *testing.T) {
	type fields struct {
		Code          code.Code
		Reason        string
		Message       string
		Metadata      map[string]string
		ForceHttpCode int32
		cause         error
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "no_metadata",
			fields: fields{
				Code: code.Code_OK,
			},
			want: []byte(`{"code":"OK","reason":"","message":""}`),
		},
		{
			name: "with_metadata",
			fields: fields{
				Code:     code.Code_OK,
				Metadata: map[string]string{"foo": "bar"},
			},
			want: []byte(`{"code":"OK","reason":"","message":"","metadata":{"foo":"bar"}}`),
		},
	}
	//goland:noinspection GoVetCopyLock
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Status: Status{
					Code:          tt.fields.Code,
					Reason:        tt.fields.Reason,
					Message:       tt.fields.Message,
					Metadata:      tt.fields.Metadata,
					ForceHttpCode: tt.fields.ForceHttpCode,
				},
				cause: tt.fields.cause,
			}
			got, err := e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	e := &Error{
		Status: Status{
			Code:    code.Code_UNKNOWN,
			Reason:  "myReason",
			Message: "myMessage",
			Metadata: map[string]string{
				"foo": "bar",
			},
			ForceHttpCode: 599,
		},
		cause: fmt.Errorf("cause"),
	}
	expected := fmt.Sprintf(`error: code = %s reason = %q message = %q metadata = %q cause = %v`,
		code.Code_UNKNOWN, "myReason", "myMessage", map[string]string{"foo": "bar"}, "cause")
	assert.Equal(t, expected, e.Error())
}

func TestError_Format(t *testing.T) {
	e := &Error{
		Status: Status{
			Code:    code.Code_UNKNOWN,
			Reason:  "myReason",
			Message: "myMessage",
			Metadata: map[string]string{
				"foo": "bar",
			},
			ForceHttpCode: 599,
		},
		cause: pkgerrors.New("cause"),
	}
	want := `error: code = UNKNOWN reason = "myReason" message = "myMessage" metadata = map["foo":"bar"] cause = cause`
	tests := []struct {
		verb string
		want string
	}{
		{
			verb: "v",
			want: want,
		},
		{
			verb: "s",
			want: want,
		},
		{
			verb: "q",
			want: fmt.Sprintf("%q", want),
		},
	}
	for _, test := range tests {
		t.Run(test.verb, func(t *testing.T) {
			got := fmt.Sprintf(fmt.Sprintf("%%%s", test.verb), e)
			assert.Equal(t, test.want, got)
		})
	}
	t.Run("+v", func(t *testing.T) {
		got := fmt.Sprintf("%+v", e)
		assert.Contains(t, got, "\n\t")
	})
}
