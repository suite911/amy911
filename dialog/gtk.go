// +build freebsd linux netbsd openbsd solaris

package dialog

import (
	"os"

	"github.com/amyadzuki/amygolib/str"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
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

func (f *GtkFrame) NewButtonRow(out *bool, kind string) {
	ok, cancel := "OK", "Cancel" // TODO: translate
	switch simp := str.Simp(kind); simp {
	case "yesno":
		ok, cancel = "Yes", "No"
	}
	hbox := gtk.NewHBox(false, 1) // TODO: what are the args for?
	bCancel := gtk.NewButtonWithLabel(cancel)
	bOk := gtk.NewButtonWithLabel(ok)
	bCancel.Clicked(func() {
		*out = false
		f.Frame.GetTopLevel().Destroy()
	})
	bOk.Clicked(func() {
		*out = true
		f.Frame.GetTopLevel().Destroy()
	})
	// ok/yes should always be to the left of cancel/no, which should be right-aligned
	// TODO: right-align these
	hbox.Add(bOk)
	hbox.Add(bCancel)
	f.Add(hbox)
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
	w.Window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})
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
