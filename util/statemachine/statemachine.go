package statemachine

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/gopherx/base/errors"
)

// T configures a single allowed transition.
type T struct {
	Current int32
	Allowed []int32
}

// Transition creates a new T object.
func Transition(current int32, allowed... int32) *T {
	return &T{current, allowed}
}

// New returns a new StateMachine.
func New(name string, strict bool, transitions ...*T) *StateMachine {
	sm := &StateMachine{name, strict, map[int32]map[int32]bool{}}

	for _, t := range transitions {
		allowed := map[int32]bool{}
		for _, a := range t.Allowed {
			allowed[a] = true
		}
		sm.config[t.Current] = allowed
	}

	return sm
}

// StateMachine implements a very simple state machine checker.
type StateMachine struct {
	name string
	strict bool
	config map[int32]map[int32]bool
}

const (
	logf = "[%s] (%d -> %d) not possible; %s"
)

func (s *StateMachine) isAllowed(cur, new int32) bool {
	if !s.strict && cur == new {
		return true
	}

	transitions, ok := s.config[cur]
	if !ok {
		glog.Errorf(logf, s.name, cur, new, "no config for current state")
		return false
	}

	config, ok := transitions[new]
	if !ok || !config {
		glog.Infof(logf, s.name, cur, new, "invalid transition")
		return false
	}

	return true
}

// Check checks if the transition is possible and returns an error if not.
func (s *StateMachine) Check(cur, new int32) error {
	if !s.isAllowed(cur, new) {
		return errors.InvalidArgument(nil, "transition not valid", fmt.Sprintf("%d -> %d", cur, new))
	}
	return nil
}