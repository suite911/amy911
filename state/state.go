package state

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

type StateCallback func(*State)

type State struct {
	fnCloseRequested  func() bool
	fnCurrent, fnNext StateCallback
	fns               map[string]StateCallback
	state             string
}

func New(fn func() bool) *State {
	return new(State).Init(fn)
}

func (s *State) Init(fn func() bool) *State {
	s.fnCloseRequested = fn
	return s
}

func (s *State) OnEnter(name string, cb StateCallback) StateBuilder {
	s.fns[str.Simp(name) + "{"] = cb
	return StateBuilder{s, name}
}

func (s *State) OnLeave(name string, cb StateCallback) StateBuilder {
	s.fns[str.Simp(name) + "}"] = cb
	return StateBuilder{s, name}
}

func (s *State) Register(name string, cb StateCallback) StateBuilder {
	s.fns[str.Simp(name)] = cb
	return StateBuilder{s, name}
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

func (s *State) SetNext(state string, ...onFail onfail.OnFail) *State {
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
	return s
}

type StateBuilder struct {
	s     *State
	state string
}

func (b StateBuilder) OnEnter(cb StateCallback) *StateBuilder {
	return b.s.OnEnter(b.state, cb)
}

func (b StateBuilder) OnLeave(cb StateCallback) *StateBuilder {
	return b.s.OnLeave(b.state, cb)
}

func (b StateBuilder) Register(cb StateCallback) *StateBuilder {
	return b.s.Register(b.state, cb)
}
