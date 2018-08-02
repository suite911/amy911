package states

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var Debug, Trace *log.Logger = nil, nil

var ErrTooManyNames =
	errors.New("State: Bad arguments to Run/RunOnce: must be () or (string)")

var MinimumSleepDuration = time.Millisecond

type State struct {
	Data interface{}

	fnCloseRequested func() bool
	fns              map[string]func(*State)
	onFail           onfail.OnFail

	editing     string
	next        string
	timestamp   time.Time
	timetosleep time.Duration
}

func New(fn func() bool) *State {
	return new(State).Init(fn)
}

func (s *State) Init(fn func() bool) *State {
	s.fnCloseRequested = fn
	s.fns = make(map[string]func(*State))
	s.timestamp = time.Now()
	return s
}

func (s *State) OnEnter(args ...interface{}) *State {
	cb := s.reg(args...)
	if Debug != nil {
		Debug.Println("OnEnter(\"" + s.editing + "\")")
	}
	s.fns[s.editing + "{"] = cb
	return s
}

func (s *State) OnFail(onFail onfail.OnFail) *State {
	s.onFail = onFail
	return s
}

func (s *State) OnFrame(args ...interface{}) *State {
	cb := s.reg(args...)
	if Debug != nil {
		Debug.Println("OnFrame(\"" + s.editing + "\")")
	}
	s.fns[s.editing] = cb
	return s
}

func (s *State) OnLeave(args ...interface{}) *State {
	cb := s.reg(args...)
	if Debug != nil {
		Debug.Println("OnLeave(\"" + s.editing + "\")")
	}
	s.fns[s.editing + "}"] = cb
	return s
}

func (s *State) Run(onFail ...onfail.OnFail) *State {
	for !s.fnCloseRequested() && len(s.next) > 0 {
		s.runOnce(onFail...)
	}
	return s
}

func (s *State) RunOnce(onFail ...onfail.OnFail) *State {
	if !s.fnCloseRequested() && len(s.next) > 0 {
		s.runOnce(onFail...)
	}
	return s
}

func (s *State) SetData(data interface{}) *State {
	s.Data = data
	return s
}

func (s *State) SetFps(fps float64) *State {
	timetosleep := float64(time.Second) / fps
	s.timetosleep = time.Duration(timetosleep)
	return s
}

func (s *State) SetNext(state string, onFail ...onfail.OnFail) *State {
	if _, ok := s.fns[str.Simp(state)]; ok {
		s.next = state
	} else {
		onfail.Fail("Cannot advance to unregistered state: \"" + state + "\"", s, nil, onFail...)
	}
	return s
}

func (s *State) Sleep() *State {
	now := time.Now()
	timestamp := s.timestamp
	s.timestamp = now
	elapsed := now.Sub(timestamp)
	remaining := elapsed - s.timetosleep
	if remaining < MinimumSleepDuration {
		remaining = MinimumSleepDuration
	}
	time.Sleep(remaining)
	return s
}

func (s *State) reg(args ...interface{}) func(*State) {
	switch len(args) {
	case 1:
		cb, ok := args[0].(func(*State))
		if ok {
			return cb
		}
	case 2:
		name, nok := args[0].(string)
		cb, cok := args[1].(func(*State))
		if nok && cok {
			s.editing = str.Simp(name)
			return cb
		}
	}
	panic(badBuilderArgs(args...))
}

func (s *State) runOnce(onFail ...onfail.OnFail) {
	state := s.next
	if Trace != nil {
		Trace.Println("Entering state: \"" + state + "\"")
	}
	main, mok := s.fns[state]
	if mok {
		enter, eok := s.fns[state + "{"]
		leave, lok := s.fns[state + "}"]
		if eok {
			enter(s)
		}
		for !s.fnCloseRequested() && s.next == state {
			main(s)
		}
		if lok {
			leave(s)
		}
	} else {
		onfail.Fail("Cannot run unregistered state: \"" + state + "\"", s, s.onFail, onFail...)
	}
	s.Sleep()
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
