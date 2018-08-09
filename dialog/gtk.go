package dialog

import (
	"os"

	"github.com/mattn/go-gtk"
)

func init() {
	setUp(
		func() {
			gtk.Init(&os.Args)
			Library = LibraryGtk{}
		},
		func() {
		},
	)
	Init()
}

type LibraryGtk struct {
	Window *gtk.Window
}

func (l *LibraryGtk) NewEntry(node Node, placeholder string, password bool) {
}

func (l *LibraryGtk) NewFrame(node Node, label string) Node {
}

func (l *LibraryGtk) NewLabel(node Node, text string) {
}

func (l *LibraryGtk) NewWindow(title string) Node {
	l.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	l.Window.SetPosition(gtk.WIN_POS_CENTER)
	l.Window.SetTitle(title)
	l.Window.SetIconName("gtk-dialog-info")
}
