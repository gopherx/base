package flag

import (
	"reflect"
	"testing"

	"github.com/gopherx/base/errors"
	"github.com/gopherx/base/errors/codes"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		args   []string
		code   codes.Code
		flags  []string
		values []string
		rem    []string
	}{
		{
			[]string{"--a"},
			codes.OK,
			[]string{"a"},
			[]string{""},
			nil,
		},
		{
			[]string{"--b", "c"},
			codes.OK,
			[]string{"b"},
			[]string{"c"},
			nil,
		},
		{
			[]string{"-d"},
			codes.OK,
			[]string{"d"},
			[]string{""},
			nil,
		},
		{
			[]string{"-e", "f"},
			codes.OK,
			[]string{"e"},
			[]string{"f"},
			nil,
		},
		{
			[]string{"--g=h"},
			codes.OK,
			[]string{"g"},
			[]string{"h"},
			nil,
		},
		{
			[]string{"-i=j"},
			codes.OK,
			[]string{"i"},
			[]string{"j"},
			nil,
		},
		{
			[]string{"k"},
			codes.InvalidArgument,
			nil,
			nil,
			nil,
		},
		{
			[]string{"--="},
			codes.InvalidArgument,
			nil,
			nil,
			nil,
		},
		{
			[]string{"--l="},
			codes.InvalidArgument,
			nil,
			nil,
			nil,
		},
		{
			[]string{"--", "doo", "foo", "bar"},
			codes.OK,
			[]string{},
			[]string{},
			[]string{"doo", "foo", "bar"},
		},
		{
			[]string{"-i", "--j", "-k='hello'", "--l", "world", "--", "doo", "foo", "bar"},
			codes.OK,
			[]string{"i", "j", "k", "l"},
			[]string{"", "", "'hello'", "world"},
			[]string{"doo", "foo", "bar"},
		},
		
	}

	for _, tc := range tests {
		flags := []string{}
		values := []string{}

		fn := func(flag, value string) error {
			flags = append(flags, flag)
			values = append(values, value)
			return nil
		}

		rem, err := Scan(tc.args, fn)
		if err != nil {
			if tc.code != errors.Code(err) {
				t.Fatal(err)
			}
			continue
		}

		if !reflect.DeepEqual(flags, tc.flags) || !reflect.DeepEqual(values, tc.values) || !reflect.DeepEqual(rem, tc.rem) {
			t.Log(tc.args)
			t.Log(flags, tc.flags)
			t.Log(values, tc.values)
			t.Log(rem, tc.rem)
			t.Error("scan failed")
		}
	}
}

func TestFDE(t *testing.T) {
	t.Log(reflect.DeepEqual([]string{}, nil))
	t.Log(reflect.DeepEqual([]string{}, []string{}))
}
