// +build freebsd linux netbsd openbsd solaris

package dialog

import (
	"os"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func init() {
	setUp(
		func() {
			gtk.Init(&os.Args)
			iLibrary = &IGtkLibrary{}
		},
		func() {
		},
	)
	Init()
}

type IGtkFrame struct {
	*gtk.Frame
}

func (f IGtkFrame) NewButtonGroup(out *int8, g *ButtonGroup) {
	hbox := gtk.NewHBox(false, 1) // TODO: what are the args for?
	for _, def := range g.Left {
		b := gtk.NewButtonWithLabel(def.Label)
		b.Clicked(func() {
			if out != nil {
				*out = def.Result
			}
			f.GetTopLevel().Destroy()
		})
		hbox.Add(b)
	}
	// TODO: right-align this one:
	b := gtk.NewButtonWithLabel(g.Right.Label)
	b.Clicked(func() {
		if out != nil {
			*out = g.Right.Result
		}
		f.GetTopLevel().Destroy()
	})
	hbox.Add(b)
	f.Add(hbox)
}

func (f IGtkFrame) NewEntry(out *string, password bool) {
	e := gtk.NewEntry()
	if out != nil {
		if placeholder := *out; len(placeholder) > 0 {
			e.SetText(placeholder)
		}
	}
	f.Add(e)
}

func (f IGtkFrame) NewLabel(text string) {
	l := gtk.NewLabel(text)
	f.Add(l)
}

type IGtkLibrary struct {
}

func (l IGtkLibrary) NewWindow(title string) Window {
	var w IGtkWindow
	w.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.SetTitle(title)
	w.SetIconName("gtk-dialog-info")
	w.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})
	return w
}

type IGtkWindow struct {
	*gtk.Window
}

func (w IGtkWindow) NewFrame(title string) Frame {
	var f IGtkFrame
	f.Frame = gtk.NewFrame(title)
	w.Add(f)
	return f
}

func (w IGtkWindow) Show(width, height int) {
	w.SetSizeRequest(width, height)
	w.ShowAll()
	gtk.Main()
}
