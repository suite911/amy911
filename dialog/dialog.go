package dialog

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var Library interface{}

type Dialog struct {
	Embed, UserData interface{}
}

func New(type string, args ...interface{}) *Dialog {
	return new(Dialog).Init(type, args...)
}

func (d *Dialog) Init(type string, args ...interface{}) *Dialog {
	switch simp := str.Simp(type); simp {
	// case "login":
	default:
		onfail.Fail("Unknown Dialog type \""+simp+"\"", d, onfail.Panic, args...)
	}
	return d
}
