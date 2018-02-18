package flag

import (
	"reflect"
	"testing"

	"github.com/gopherx/base/errors"
	"github.com/gopherx/base/errors/codes"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		args  []string
		code  codes.Code
		flags []Flag
		rem   []string
	}{
		{
			[]string{"--a"},
			codes.OK,
			[]Flag{{"a", "", "--", ""}},
			nil,
		},
		{
			[]string{"--b", "c"},
			codes.OK,
			[]Flag{{"b", "c", "--", " "}},
			nil,
		},
		{
			[]string{"-d"},
			codes.OK,
			[]Flag{{"d", "", "-", ""}},
			nil,
		},
		{
			[]string{"-e", "f"},
			codes.OK,
			[]Flag{{"e", "f", "-", " "}},
			nil,
		},
		{
			[]string{"--g=h"},
			codes.OK,
			[]Flag{{"g", "h", "--", "="}},
			nil,
		},
		{
			[]string{"-i=j"},
			codes.OK,
			[]Flag{{"i", "j", "-", "="}},
			nil,
		},
		{
			[]string{"k"},
			codes.InvalidArgument,
			nil,
			nil,
		},
		{
			[]string{"--="},
			codes.InvalidArgument,
			nil,
			nil,
		},
		{
			[]string{"--l="},
			codes.InvalidArgument,
			nil,
			nil,
		},
		{
			[]string{"--", "doo", "foo", "bar"},
			codes.OK,
			[]Flag{},
			[]string{"doo", "foo", "bar"},
		},
		{
			[]string{"-i", "--j", "-k='hello'", "--l", "world", "--", "doo", "foo", "bar"},
			codes.OK,
			[]Flag{
				{"i", "", "-", ""},
				{"j", "", "--", ""},
				{"k", "'hello'", "-", "="},
				{"l", "world", "--", " "},
			},
			[]string{"doo", "foo", "bar"},
		},
		{
			[]string{"------m", "doo", "--", "foo", "bar"},
			codes.OK,
			[]Flag{{"m", "doo", "------", " "}},
			[]string{"foo", "bar"},
		},
	}

	for _, tc := range tests {
		flags := []Flag{}

		fn := func(f Flag) error {
			flags = append(flags, f)
			return nil
		}

		rem, err := Scan(tc.args, fn)
		if err != nil {
			if tc.code != errors.Code(err) {
				t.Fatal(err)
			}
			continue
		}

		if !reflect.DeepEqual(flags, tc.flags) {
			t.Log(tc)
			t.Log(tc.flags)
			t.Log(flags)
			t.Log(rem, tc.rem)
			t.Error("scan failed")
		}
	}
}

func TestFDE(t *testing.T) {
	t.Log(reflect.DeepEqual([]string{}, nil))
	t.Log(reflect.DeepEqual([]string{}, []string{}))
}
