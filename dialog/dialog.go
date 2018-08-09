package dialog

import (
	"sync"

	"github.com/amyadzuki/amygolib/onfail"
	"github.com/amyadzuki/amygolib/str"
)

var (
	OverrideLogIn func(*Dialog, []interface{})
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
			w := Library.NewWindow("Log in") // TODO: translate
			f := w.NewFrame("Account") // TODO: translate
			f.NewLabel("E-mail address:") // TODO: translate
			f.NewEntry("", false)
			f.NewLabel("Password:") // TODO: translate
			f.NewEntry("", true)
			w.Show(576, 324)
		}
	default:
		onfail.Fail("Unknown Dialog kind \""+simp+"\"", d, onfail.Panic, args...)
	}
	return d
}
