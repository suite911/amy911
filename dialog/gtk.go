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
			iLibrary = GtkLibrary{}
		},
		func() {
		},
	)
	Init()
}

type GtkFrame struct {
	Frame *gtk.Frame
}

func (f *GtkFrame) NewButtonGroup(out *int8, g *ButtonGroup) {
	hbox := gtk.NewHBox(false, 1) // TODO: what are the args for?
	for _, def := range g.Left {
		b := gtk.NewButtonWithLabel(def.Label)
		b.Clicked(func() {
			if out != nil {
				*out = def.Result
			}
			f.Frame.GetTopLevel().Destroy()
		})
		hbox.Add(b)
	}
	// TODO: right-align this one:
	b := gtk.NewButtonWithLabel(g.Right.Label)
	b.Clicked(func() {
		if out != nil {
			*out = g.Right.Result
		}
		f.Frame.GetTopLevel().Destroy()
	})
	hbox.Add(b)
	f.Add(hbox)
}

func (f *GtkFrame) NewEntry(out *string, password bool) {
	e := gtk.NewEntry()
	if out != nil {
		if placeholder = *out; len(placeholder) > 0 {
			e.SetText(placeholder)
		}
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
