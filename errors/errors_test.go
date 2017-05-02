package errors

import (
	"errors"
	"testing"

	"github.com/gopherx/base/errors/codes"
	"strings"
	"fmt"
)

// TestConstructors checks that we got a constructor func for every code in the codes package.
func TestConstructors(t *testing.T) {
	tests := map[codes.Code]ErrorFunc{
		codes.Canceled:           Canceled,
		codes.Unknown:            Unknown,
		codes.InvalidArgument:    InvalidArgument,
		codes.DeadlineExceeded:   DeadlineExceeded,
		codes.NotFound:           NotFound,
		codes.AlreadyExists:      AlreadyExists,
		codes.PermissionDenied:   PermissionDenied,
		codes.Unauthenticated:    Unauthenticated,
		codes.ResourceExhausted:  ResourceExhausted,
		codes.FailedPrecondition: FailedPrecondition,
		codes.Aborted:            Aborted,
		codes.OutOfRange:         OutOfRange,
		codes.Unimplemented:      Unimplemented,
		codes.Internal:           Internal,
		codes.Unavailable:        Unavailable,
		codes.DataLoss:           DataLoss,
	}

	cause := errors.New("the cause of the problem")
	mid := Internal(cause, "waaat")
	desc := "oops an error occured"
	a0 := 1
	a1 := 0.5
	a2 := "wrong"
	for code, ctor := range tests {
		err := ctor(mid, desc, a0, a1, a2)
		if Code(err) != code {
			t.Error(Code(err), code)
		}

		if Cause(err) != mid {
			t.Error(Cause(err), mid)
		}

		// Check the error text! Note that the text is not static and will change with
		// the environment and therefore we can't simply compare against static strings.
		eparts := strings.Split(err.Error(), "\n")
		checkDescLine := func(i int, indent string, code codes.Code, desc string) {
			werr := indent + code.String() + "] " + desc
			if eparts[i] != werr {
				t.Error(fmt.Sprintf("%q %q", eparts[i], werr))
			}
		}

		//..."check" some of the callstack
		test := "gopherx/base/errors.TestConstructors"
		file := "errors_test.go:"
		checkStackLine := func(i int, indent string) {
			if strings.Index(eparts[i], indent) != 0 {
				goto fail
			}

			if !strings.Contains(eparts[i], file) {
				goto fail
			}

			if !strings.Contains(eparts[i], test) {
				goto fail
			}

			return

			fail:
				t.Error(eparts[i])
		}

		checkDescLine(0, "", code, desc)
		checkStackLine(1, "")
		checkDescLine(3, "  ", codes.Internal, "waaat")
		checkStackLine(4, "  ")

		//for _, p := range eparts {
		//	t.Log(p)
		//}
	}
}

func recCallers(n int) []uintptr {
	if n == 0 {
		return callers(0)
	}

	return recCallers(n - 1)
}

func TestCallers(t *testing.T) {
	cs := recCallers(10)
	if len(cs) < 10 {
		t.Fatal(len(cs))
	}
}
