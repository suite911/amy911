// +build freebsd linux netbsd openbsd solaris

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

type GtkFrame struct {
	Frame *gtk.Frame
}

func (f *GtkFrame) NewEntry(out *string, placeholder string, password bool) {
	e := gtk.NewEntry()
	if len(placeholder) > 0 {
		e.SetText(placeholder)
	}
	f.Add(e)
}

func (f *GtkFrame) NewLabel(text string) {
	l := gtk.NewLabel(text)
	f.Add(l)
}

type GtkLibrary struct {
}

func (l *GtkLibrary) NewWindow(title string) Window {
	var w GtkWindow
	w.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	w.Window.SetPosition(gtk.WIN_POS_CENTER)
	w.Window.SetTitle(title)
	w.Window.SetIconName("gtk-dialog-info")
	return w
}

type GtkWindow struct {
	Window *gtk.Window
}

func (w *GtkWindow) NewFrame(title string) Window {
	var f GtkFrame
	f.Frame = gtk.NewFrame(title)
	w.Add(f)
	return f
}

func (w *GtkWindow) Show(width, height int) {
	w.SetSizeRequest(width, height)
	w.ShowAll()
	gtk.Main()
}
