package states

import (
	"errors"

	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

type Callback func(*State)
var ErrBuilderArgs =
	errors.New("State: Bad arguments to builder method: must be (Callback) or (string, Callback)")

type State struct {
	Data interface{}

	fnCloseRequested func() bool
	fns              map[string]Callback
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
}

func (s *State) SetData(data interface{}) *State {
	s.Data = data
	return s
}

func (s *State) SetNext(state string, onFail ...onfail.OnFail) *State {
	if _, ok := s.fns[str.Simp(state)]; ok {
		s.sNext = state
	} else {
		var failFunc onfail.OnFail = onfail.Panic
		if len(onFail) > 0 {
			failFunc = onFail[0]
		}
		failFunc.Fail(errors.New("Unregistered state: \"" + state + "\""))
	}
	return s
}

type registrationBuilder struct {
	s     *State
	state string
}

func (b registrationBuilder) OnEnter(args ...interface{}) registrationBuilder {
	a := b.reg(args...)
	return b.s.OnEnter(a.state, a.cb)
}

func (b registrationBuilder) OnLeave(args ...interface{}) registrationBuilder {
	a := b.reg(args...)
	return b.s.OnLeave(a.state, a.cb)
}

func (b registrationBuilder) Register(args ...interface{}) registrationBuilder {
	a := b.reg(args...)
	return b.s.Register(a.state, a.cb)
}

func (b registrationBuilder) reg(args ...interface{}) registrationBuilderArgs {
	if len(args) == 1 || len(args) == 2 {
		switch args[0].(type) {
		case Callback:
			if len(args) == 1 {
				return registrationBuilderArgs{args[0].(Callback), b.state}
			}
		case string:
			if len(args) == 2 {
				switch args[1].(type) {
				case Callback:
					return registrationBuilderArgs{args[1].(Callback), args[0].(string)}
				}
			}
		}
	}
	panic(ErrBuilderArgs)
}

type registrationBuilderArgs struct {
	cb    Callback
	state string
}
