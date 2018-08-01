package state

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

type State struct {
	fnCloseRequested  func() bool
	fnCurrent, fnNext func(*State)
	fns               map[string]func(*State)
	state             string
}

func New(fn func() bool) *State {
	s := new(State)
	s.Init(fn)
	return s
}

func (s *State) Init(fn func() bool) {
	s.fnCloseRequested = fn
}

func (s *State) OnEnter(name string, cb func(*State)) {
	s.fns[str.Simp(name) + "{"] = cb
}

func (s *State) OnLeave(name string, cb func(*State)) {
	s.fns[str.Simp(name) + "}"] = cb
}

func (s *State) Register(name string, cb func(*State)) {
	s.fns[str.Simp(name)] = cb
}

func (s *State) Run() {
	s.fnCurrent = s.fnNext
	enter, eok := s.fns[s.state + "{"]
	leave, lok := s.fns[s.state + "}"]
	if eok {
		enter(s)
	}
	for !s.fnCloseRequested() && s.fnNext == s.fnCurrent {
		s.fnCurrent(s)
	}
	if lok {
		leave(s)
	}
}

func (s *State) SetNext(state string, ...onFail onfail.OnFail) {
	if fn, ok := s.fns[state], ok {
		s.state = state
		s.fnNext = fn
	} else {
		failFunc := onfail.Panic
		if len(onFail) > 0 {
			failFunc = onFail[0]
		}
		failFunc("Unregistered state: \"" + state + "\"")
	}
}
