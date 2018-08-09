package dialog

import "github.com/amy911/amy911/str"

// From <https://docs.microsoft.com/en-us/windows/desktop/uxguide/top-violations>
//     "Present the commit buttons in the following order:"
//         "OK/[Do it]/Yes"
//         "[Don't do it]/No"
//         "Cancel"
//         "Apply (if present)"

const (
	Cancel int8 = iota
	Yes
	No
)
const LogIn = Yes // Yes, I have an account
const Register = No // No, I do not have an account

type ButtonDef struct {
	Label  string
	Result int8
}

type ButtonGroup struct {
	Left []ButtonDef
	Right ButtonDef
}

func NewButtonGroup(kind string) *ButtonGroup {
	return new(ButtonGroup).Init(kind)
}

func (g *ButtonGroup) Init(kind string) *ButtonGroup {
	ok := ButtonDef{"OK", Yes} // TODO: translate
	yes := ButtonDef{"Yes", Yes} // TODO: translate
	no := ButtonDef{"No", No} // TODO: translate
	close_ := ButtonDef{"Close", Cancel} // TODO: translate
	cancel := ButtonDef{"Cancel", Cancel} // TODO: translate
	login := ButtonDef{"Log In", LogIn} // TODO: translate
	register := ButtonDef{"Register", Register} // TODO: translate

	g.Left = []ButtonDef{ok}
	g.Right = cancel
	switch simp := str.Simp(kind); simp {
	case "ok":
		g.Left = []ButtonDef{}
		g.Right = ok
	case "close", "error", "problem", "warning":
		g.Left = []ButtonDef{}
		g.Right = close_
	case "login":
		g.Left = []ButtonDef{login, register}
		g.Right = cancel
	case "yesno", "yesnocancel":
		g.Left = []ButtonDef{yes, no}
		g.Right = cancel
	}
	return g
}
