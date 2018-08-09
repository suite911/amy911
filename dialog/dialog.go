package dialog

import (
	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var Library interface{}

type Dialog struct {
	Embed, UserData interface{}
}

func New(type string, onFail ...onfail.OnFail) *Dialog {
	return new(Dialog).Init(type, onFail...)
}

func (d *Dialog) Init(type string, onFail ...onfail.OnFail) *Dialog {
	switch simp := str.Simp(type); simp {
	// case "login":
	default:
		onfail.Fail("Unknown Dialog type \""+simp+"\"", d, onfail.Panic, onFail...)
	}
	return d
}
