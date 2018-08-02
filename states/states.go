package states

import (
	"errors"
	"fmt"

	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var ErrTooManyNames =
	errors.New("State: Bad arguments to Run: must be () or (string)")

type State struct {
	Data interface{}

	fnCloseRequested func() bool
	fns              map[string]func(*State)
	editing          string
	sCurrent, sNext  string
	state            string
}

func New(fn func() bool) *State {
	return new(State).Init(fn)
}

func (s *State) Init(fn func() bool) *State {
	s.fnCloseRequested = fn
	return s
}

func (s *State) OnEnter(args ...interface{}) *State {
	s.fns[s.editing + "{"] = s.reg(args...)
	return s
}

func (s *State) OnLeave(args ...interface{}) *State {
	s.fns[s.editing + "}"] = s.reg(args...)
	return s
}

func (s *State) Register(args ...interface{}) *State {
	s.fns[s.editing] = s.reg(args...)
	return s
}

func (s *State) Run(name ...string) *State {
	switch len(name) {
	case 0:
	case 1:
		s.sNext = name[0]
	default:
		panic(ErrTooManyNames)
	}
	s.sCurrent = s.sNext
	main, mok := s.fns[s.state]
	if mok {
		enter, eok := s.fns[s.state + "{"]
		leave, lok := s.fns[s.state + "}"]
		if eok {
			enter(s)
		}
		for !s.fnCloseRequested() && s.sNext == s.sCurrent {
			main(s)
		}
		if lok {
			leave(s)
		}
	}
	return s
}

func (s *State) SetData(data interface{}) *State {
	s.Data = data
	return s
}

func (s *State) SetNext(state string, onFail ...onfail.OnFail) *State {
	if _, ok := s.fns[str.Simp(state)]; ok {
		s.sNext = state
	} else {
		onfail.Fail("Unregistered state: \"" + state + "\"", s, nil, onFail...)
	}
	return s
}

func (s *State) reg(args ...interface{}) func(*State) {
	if len(args) == 1 || len(args) == 2 {
		switch args[0].(type) {
		case func(*State):
			if len(args) == 1 {
				return args[0].(func(*State))
			}
		case string:
			if len(args) == 2 {
				switch args[1].(type) {
				case func(*State):
					s.editing = args[0].(string)
					return args[1].(func(*State))
				}
			}
		}
	}
	panic(badBuilderArgs(args...))
}

func badBuilderArgs(args ...interface{}) error {
	msg := "State: Bad arguments to builder method: must be (func(*State)) or (string, func(*State))\nHave: ("
	for aid, arg := range args {
		if aid != 0 {
			msg += ", "
		}
		msg += fmt.Sprintf("%T", arg)
	}
	msg += ")"
	return errors.New(msg)
}
