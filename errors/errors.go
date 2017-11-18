// Package errors provides a very opionionated interface for working with errors; the opions are:
// - the error path is the least tested and excercised path in the entire system
// - errors are not free; don't build interfaces that return errors as part of normal excecution
// - errors should capture as much information as possible to simplify debugging
// - use minimal APIs to minimize the risk of bad/missing args etc
//
// Using these opinions as truth:
// - all root errors always capture the stacktrace of the caller (root error, first error from this package)
// - no fmt.Errorf(...) like formatting APIs are available
//
// The codes used by this package are copied from the grpc project. (https://github.com/grpc/grpc-go)
package errors

import (
	"fmt"
	"github.com/gopherx/base/errors/codes"
	"runtime"
)

type eee struct {
	code    codes.Code
	cause   error
	desc    string
	args    []interface{}
	callers []uintptr
}

func callers(skip int) []uintptr {
	const initial = 64
	size := initial
	for {
		cs := make([]uintptr, size)
		n := runtime.Callers(skip, cs)

		if n < len(cs) {
			return cs[:n]
		}

		size = size * 2
	}
}

func newEee(code codes.Code, cause error, desc string, args []interface{}) *eee {
	return &eee{code, cause, desc, args, callers(4)}
}

// Error implements the error interface.
func (e *eee) Error() string {
	return fmt.Sprint(e)
}

var (
	// FormatError formats the error into the buffer. Override to change error output.
	FormatError = func(s fmt.State,
		c rune,
		indent string,
		err error,
		code codes.Code,
		desc string,
		args []interface{},
		callers []uintptr) {

		//...is the error from another package?
		if err != nil {
			fmt.Fprint(s, indent, "error] ", err)
			return
		}

		fmt.Fprint(s,
			indent,
			code.String(), "] ",
			desc,
			" args:",
			args,
		)

		frames := runtime.CallersFrames(callers)
		for {
			frame, ok := frames.Next()
			if !ok {
				return
			}
			s.Write(newLine)
			fmt.Fprint(s, indent, frame.File, ":", frame.Line, " ", frame.Func.Name())
		}
	}

	// Array to use for indents of causing errors. Last indent string is used once max is reached.
	Indents = []string{"", "  ", "    ", "      ", "        "}

	newLine = []byte{'\n'}
	empty   = []byte{}
)

// Format implements the fmt.Formatter interface.
func (e *eee) Format(s fmt.State, c rune) {
	cur := e
	cnt := 0
	indent := ""
	sep := empty

	updateIndent := func() {
		if cnt < len(Indents) {
			indent = Indents[cnt]
		}
	}

	for {
		s.Write(sep)
		updateIndent()

		FormatError(s, c, indent, nil, cur.code, cur.desc, cur.args, cur.callers)
		// TODO(d): should only print newline when needed and not always

		cnt++

		nxt, ok := cur.cause.(*eee)
		if ok {
			// more eee:s to format
			cur = nxt
			sep = newLine
			continue
		}

		if cur.cause == nil {
			return
		}

		// Root cause is not a eee error.
		s.Write(sep)
		cnt++
		updateIndent()
		FormatError(s, c, indent, cur.cause, codes.OK, "", nil, nil)
		return
	}
}

// Code returns the code of the error. Returns codes.Unknown if not an error from this package.
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	e, ok := err.(*eee)
	if !ok {
		return codes.Unknown
	}

	return e.code
}

// Cause returns the cause of the error (or nil if not set or an error not created by this package).
func Cause(err error) error {
	e, ok := err.(*eee)
	if !ok {
		return nil
	}
	return e.cause
}

// ErrorFunc is a function that creates an error.
type ErrorFunc func(cause error, desc string, args ...interface{}) error

// Canceled returns a new codes.Canceled error.
func Canceled(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Canceled, cause, desc, args)
}

// Unknown returns a new codes.Unknown error.
func Unknown(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Unknown, cause, desc, args)
}

// InvalidArgument returns a new codes.InvalidArgument error.
func InvalidArgument(cause error, desc string, args ...interface{}) error {
	return newEee(codes.InvalidArgument, cause, desc, args)
}

// DeadlineExceeded returns a new codes.DeadlineExceeded error.
func DeadlineExceeded(cause error, desc string, args ...interface{}) error {
	return newEee(codes.DeadlineExceeded, cause, desc, args)
}

// NotFound returns a new codes.NotFound error.
func NotFound(cause error, desc string, args ...interface{}) error {
	return newEee(codes.NotFound, cause, desc, args)
}

// AlreadyExists returns a new codes.AlreadyExists error.
func AlreadyExists(cause error, desc string, args ...interface{}) error {
	return newEee(codes.AlreadyExists, cause, desc, args)
}

// PermissionDenied returns a new codes.PermissionDenied error.
func PermissionDenied(cause error, desc string, args ...interface{}) error {
	return newEee(codes.PermissionDenied, cause, desc, args)
}

// Unauthenticated returns a new codes.Unauthenticated error.
func Unauthenticated(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Unauthenticated, cause, desc, args)
}

// ResourceExhausted returns a new codes.ResourceExhausted error.
func ResourceExhausted(cause error, desc string, args ...interface{}) error {
	return newEee(codes.ResourceExhausted, cause, desc, args)
}

// FailedPrecondition returns a new codes.FailedPrecondition error.
func FailedPrecondition(cause error, desc string, args ...interface{}) error {
	return newEee(codes.FailedPrecondition, cause, desc, args)
}

// Aborted returns a new codes.Aborted error.
func Aborted(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Aborted, cause, desc, args)
}

// OutOfRange returns a new codes.OutOfRange error.
func OutOfRange(cause error, desc string, args ...interface{}) error {
	return newEee(codes.OutOfRange, cause, desc, args)
}

// Unimplemented returns a new codes.Unimplemented error.
func Unimplemented(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Unimplemented, cause, desc, args)
}

// Internal returns a new codes.Internal error.
func Internal(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Internal, cause, desc, args)
}

// Unavailable returns a new codes.Unavailable error.
func Unavailable(cause error, desc string, args ...interface{}) error {
	return newEee(codes.Unavailable, cause, desc, args)
}

// DataLoss returns a new codes.DataLoss error.
func DataLoss(cause error, desc string, args ...interface{}) error {
	return newEee(codes.DataLoss, cause, desc, args)
}
