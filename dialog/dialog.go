package dialog

import (
	"sync"

	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var (
	Init = func() {}
	Quit = func() {}

	Library ILibrary

	OverrideLogIn func(*Dialog, []interface{})

	mutex sync.Mutex
)

type Dialog struct {
	Embed, UserData interface{}
}

func New(kind string, args ...interface{}) *Dialog {
	return new(Dialog).Init(kind, args...)
}

func (d *Dialog) Init(kind string, args ...interface{}) *Dialog {
	if Library == nil {
		onfail.Fail("Dialog Library was not loaded or was unloaded", d, onfail.Panic, args...)
		return d
	}
	switch simp := str.Simp(kind); simp {
	case "login":
		if OverrideLogIn != nil {
			OverrideLogIn(d, args)
		} else {
		}
	default:
		onfail.Fail("Unknown Dialog kind \""+simp+"\"", d, onfail.Panic, args...)
	}
	return d
}

type ILibrary interface {
}
