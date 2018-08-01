package states

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

type Callback func(*State)

type State struct {
	Data interface{}

	fnCloseRequested  func() bool
	fnCurrent, fnNext Callback
	fns               map[string]Callback
	state             string
}

func New(fn func() bool) *State {
	return new(State).Init(fn)
}

func (s *State) Init(fn func() bool) *State {
	s.fnCloseRequested = fn
	return s
}

func (s *State) OnEnter(name string, cb Callback) registrationBuilder {
	s.fns[str.Simp(name) + "{"] = cb
	return registrationBuilder{s, name}
}

func (s *State) OnLeave(name string, cb Callback) registrationBuilder {
	s.fns[str.Simp(name) + "}"] = cb
	return registrationBuilder{s, name}
}

func (s *State) Register(name string, cb Callback) registrationBuilder {
	s.fns[str.Simp(name)] = cb
	return registrationBuilder{s, name}
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

func (s *State) SetData(data interface{}) *State {
	s.Data = data
	return s
}

func (s *State) SetNext(state string, onFail ...onfail.OnFail) *State {
	if fn, ok := s.fns[str.Simp(state)]; ok {
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

type registrationBuilder struct {
	s     *State
	state string
}

func (b registrationBuilder) OnEnter(args ...interface{}) *registrationBuilder {
	a := b.reg(args...)
	return b.s.OnEnter(a.state, a.cb)
}

func (b registrationBuilder) OnLeave(args ...interface{}) *registrationBuilder {
	a := b.reg(args...)
	return b.s.OnLeave(a.state, a.cb)
}

func (b registrationBuilder) Register(args ...interface{}) *registrationBuilder {
	a := b.reg(args...)
	return b.s.Register(a.state, a.cb)
}

func (b registrationBuilder) reg(args ...interface{}) registrationBuilderArgs {
	switch {
	case len(args) == 1 && args[0].(type) == Callback:
		return registrationBuilderArgs{args[0].(Callback), b.state}
	case len(args) == 2 && args[0].(type) == string && args[1].(type) == Callback:
		return registrationBuilderArgs{args[1].(Callback), args[0].(string)}
	default:
		panic("State: Bad arguments to builder method: must be (Callback) or (string, Callback)")
	}
}

type registrationBuilderArgs struct {
	cb    Callback
	state string
}
