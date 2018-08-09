package dialog

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var (
	Library interface{}

	OverrideLogIn func(*Dialog, []interface{})
)

type Dialog struct {
	Embed, UserData interface{}
}

func New(type string, args ...interface{}) *Dialog {
	return new(Dialog).Init(type, args...)
}

func (d *Dialog) Init(type string, args ...interface{}) *Dialog {
	if Library == nil {
		onfail.Fail("Dialog Library was not loaded or was unloaded", d, onfail.Panic, args...)
		return d
	}
	switch simp := str.Simp(type); simp {
	case "login":
		if OverrideLogIn != nil {
			OverrideLogIn(d, args)
		} else {
		}
	default:
		onfail.Fail("Unknown Dialog type \""+simp+"\"", d, onfail.Panic, args...)
	}
	return d
}
