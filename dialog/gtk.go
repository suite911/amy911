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

func (l *LibraryGtk) NewWindow(title string) {
	l.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	l.Window.SetPosition(gtk.WIN_POS_CENTER)
	l.Window.SetTitle(title)
	l.Window.SetIconName("gtk-dialog-info")
}
