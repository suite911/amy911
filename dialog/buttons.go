package dialog

import "github.com/amyadzuki/amygolib/str"

// From <https://docs.microsoft.com/en-us/windows/desktop/uxguide/top-violations>
//     "Present the commit buttons in the following order:"
//         "OK/[Do it]/Yes"
//         "[Don't do it]/No"
//         "Cancel"
//         "Apply (if present)"

type Result int {
	Cancel = iota
	Yes
	No
}

type ButtonGroup struct {
	Left []string
	Right string
}

func NewButtonGroup(kind string) *ButtonGroup {
	return new(ButtonGroup).Init(kind)
}

func (g *ButtonGroup) Init(kind string) *ButtonGroup {
	g.Left = []string{"OK"}
	g.Right = "Cancel"
	switch simp := str.Simp(kind); simp {
	case "ok":
		g.Left = []string{}
		g.Right = "OK" // TODO: translate
	case "close", "error", "problem", "warning":
		g.Left = []string{}
		g.Right = "Close" // TODO: translate
	case "yesno", "yesnocancel":
		g.Left = []string{"Yes", "No"} // TODO: translate
		g.Right = "Cancel" // TODO: translate
	}
	return g
}
