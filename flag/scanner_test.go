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
		flags []Spec
		rem   []string
	}{
		{
			[]string{"--a"},
			codes.OK,
			[]Spec{{"a", "", "--", ""}},
			nil,
		},
		{
			[]string{"--b", "c"},
			codes.OK,
			[]Spec{{"b", "c", "--", " "}},
			nil,
		},
		{
			[]string{"-d"},
			codes.OK,
			[]Spec{{"d", "", "-", ""}},
			nil,
		},
		{
			[]string{"-e", "f"},
			codes.OK,
			[]Spec{{"e", "f", "-", " "}},
			nil,
		},
		{
			[]string{"--g=h"},
			codes.OK,
			[]Spec{{"g", "h", "--", "="}},
			nil,
		},
		{
			[]string{"-i=j"},
			codes.OK,
			[]Spec{{"i", "j", "-", "="}},
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
			[]Spec{},
			[]string{"doo", "foo", "bar"},
		},
		{
			[]string{"-i", "--j", "-k='hello'", "--l", "world", "--", "doo", "foo", "bar"},
			codes.OK,
			[]Spec{
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
			[]Spec{{"m", "doo", "------", " "}},
			[]string{"foo", "bar"},
		},
		{
			[]string{},
			codes.OK,
			[]Spec{},
			nil,
		},
		{
			[]string{"--x", "hello", "world", "--", "tree"},
			codes.OK,
			[]Spec{{"x", "hello", "--", " "}, {"world", "", "", ""}},
			[]string{"tree"},
		},
		{
			[]string{"x", "--", "tree"},
			codes.OK,
			[]Spec{{"x", "", "", ""}},
			[]string{"tree"},
		},
		{
			[]string{"hello", "--", "tree"},
			codes.OK,
			[]Spec{{"hello", "", "", ""}},
			[]string{"tree"},
		},
		{
			[]string{"hello", "world", "--", "tree"},
			codes.OK,
			[]Spec{{"hello", "", "", ""}, {"world", "", "", ""}},
			[]string{"tree"},
		},
		{
			[]string{"hello=foo", "world=bar", "car", "--", "tree"},
			codes.OK,
			[]Spec{{"hello", "foo", "", "="}, {"world", "bar", "", "="}, {"car", "", "", ""}},
			[]string{"tree"},
		},
	}

	for _, tc := range tests {
		t.Log(tc.args)

		flags := []Spec{}

		fn := func(f Spec) error {
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

		if !reflect.DeepEqual(flags, tc.flags) || !reflect.DeepEqual(rem, tc.rem) {
			t.Log("case", tc)
			t.Log("g.flags", flags, flags == nil)
			t.Log("w.flags", tc.flags, tc.flags == nil)
			t.Log("g.rem", rem, rem == nil)
			t.Log("w.rem", tc.rem, tc.rem == nil)

			t.Fatal("scan failed")
		}

		t.Log("OK")
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		txt string
		h  string
		n  string
	}{
		{"", "", ""},
		{"x", "", "x"},
		{"--x", "--", "x"},
		{"---x", "---", "x"},
		{"-x", "-", "x"},
		{"-", "", "-"},
		{"x-", "", "x-"},
	}

	for _, tc := range tests {
		t.Log(tc)
		
		h, n := split(tc.txt)
		if h != tc.h || n != tc.n {
			t.Log("g.header", h)
			t.Log("w.header", tc.h)
			t.Log("g.name", n)
			t.Log("w.name", tc.n)
			t.Fatal("split failed")
		}

		t.Log("OK")
	}
}

func TestCut(t *testing.T) {
	tests := []struct {
		l []string
		h string
		t []string
	} {
		{[]string{}, "", nil},
		{[]string{"a"}, "a", nil},
		{[]string{"b", "c", "d"}, "b", []string{"c", "d"}},
	}

	for _, tc := range tests {
		t.Log(tc)
		
		head, tail := cut(tc.l)
		if head != tc.h || !reflect.DeepEqual(tail, tc.t) {
			t.Log("g.head", head)
			t.Log("w.head", tc.h)
			t.Log("g.tail", tail)
			t.Log("w.tail", tc.t)
			t.Fatal()
		}

		t.Log("OK")
	}
}
