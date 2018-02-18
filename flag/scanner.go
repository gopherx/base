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
// If the flag have no value then the value string is empty.
func Scan(args []string, fn FlagFunc) ([]string, error) {
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

func cut(s []string) (string, []string) {
	h := s[0]
	var t []string
	if len(s) > 1 {
		t = s[1:]
	}
	return h, t
}

func split(f string) (string, string) {
	nameAt := 0
	for _, c := range f {
		if c == '-' {
			nameAt++
			continue
		}

		break
	}

	return f[:nameAt], f[nameAt:]
}

func scanFirstFlag(args []string, fn FlagFunc) ([]string, bool, error) {
	head, tail := cut(args)
	if len(head) < 2 {
		//...too short; can't be a flag.
		return nil, false, errors.InvalidArgument(nil, "first argument too short; must be >= 2", args)
	}

	header, name := split(head)
	value := ""
	if len(name) == 0 {
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

	if len(value) == 0 && len(tail) > 0 {
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
