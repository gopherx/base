package flag

import (
	"strings"

	"github.com/gopherx/base/errors"
)

// Spec holds all data from a parsed flag.
type Spec struct {
	Name      string
	Value     string
	Header    string
	Separator string
}

// FlagFunc is called by Scan for every found flag.
type FlagFunc func(f Spec) error

// Scan scans Args for flags and issues callbacks when a new flag is found.
// All flags are assumed to start with '-' or '--' or actually any '-'.
// They may be followed by a value that may or may not be separated by '='.
// If the flag have no value then the Value field is empty.
// If the flag have no header then only flags with '=' will have a value.
func Scan(args []string, fn FlagFunc) ([]string, error) {
	if len(args) == 0 {
		return nil, nil
	}

	next := args
	for {
		rem, term, err := scanFirstFlag(next, fn)
		if err != nil {
			return nil, err
		}

		if term || len(rem) == 0 {
			return rem, nil
		}

		next = rem
	}
}

// cut cuts the head of the list and returns the head and tail.
func cut(s []string) (string, []string) {
	if len(s) == 0 {
		return "", nil
	}

	h := s[0]
	var t []string
	if len(s) > 1 {
		t = s[1:]
	}
	return h, t
}

// split splits a flag '--N' into '--' and 'N'. The leading dashes are considered
// to be a header iff they start at index zero and end at least one char before
// the end of the string. A string with only dashes will return the dashes as name.
func split(f string) (string, string) {
	dashCnt := 0
	for _, c := range f {
		if c == '-' {
			dashCnt++
			continue
		}

		break
	}

	if dashCnt == 0 || dashCnt == len(f) {
		return "", f
	}

	return f[:dashCnt], f[dashCnt:]
}

func scanFirstFlag(args []string, fn FlagFunc) ([]string, bool, error) {
	ff, tail := cut(args)
	if len(ff) == 0 {
		return nil, false, nil
	}

	header, name := split(ff)
	value := ""
	if len(header) == 0 && name == "--" {
		//...terminator '--' found
		return tail, true, nil
	}

	separator := "="
	eqAt := strings.IndexByte(name, '=')
	if eqAt == 0 {
		return nil, false, errors.InvalidArgument(nil, "malformed flag found", args)
	}

	if eqAt > 0 {
		value = name[eqAt+1:]
		name = name[:eqAt]
		if len(value) == 0 {
			return nil, false, errors.InvalidArgument(nil, "malformed flag found", args)
		}
	}

	if len(value) == 0 && len(tail) > 0 && len(header) > 0 {
		peek, tmp := cut(tail)
		if len(peek) > 0 && peek[0] != '-' {
			value = peek
			tail = tmp
			separator = " "
		}
	}

	if len(value) == 0 {
		separator = ""
	}

	err := fn(Spec{name, value, header, separator})
	return tail, false, err
}
