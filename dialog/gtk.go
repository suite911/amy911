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
			iLibrary = &GtkLibrary{}
		},
		func() {
		},
	)
	Init()
}

type GtkFrame struct {
	Frame *gtk.Frame
	*gtk.VBox
}

func (f GtkFrame) NewEntry(out *string, password bool) {
	e := gtk.NewEntry()
	if out != nil {
		if placeholder := *out; len(placeholder) > 0 {
			e.SetText(placeholder)
		}
	}
	f.Add(e)
}

func (f GtkFrame) NewLabel(text string) {
	l := gtk.NewLabel(text)
	f.Add(l)
}

type GtkLibrary struct {
}

func (l GtkLibrary) NewWindow(title string) Window {
	var w GtkWindow
	w.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	w.Window.SetPosition(gtk.WIN_POS_CENTER)
	w.Window.SetTitle(title)
	w.Window.SetIconName("gtk-dialog-info")
	w.Window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})
	w.VBox = gtk.NewVBox(true, 1) // (homogeneous bool, spacing int)
	w.Window.Add(w.VBox)
	return w
}

type GtkWindow struct {
	Window *gtk.Window
	*gtk.VBox
}

func (w GtkWindow) NewButtonGroup(out *int8, g *ButtonGroup) {
	hbox := gtk.NewHBox(true, 1) // (homogeneous bool, spacing int)
	for _, def := range g.Left {
		b := gtk.NewButtonWithLabel(def.Label)
		b.Clicked(func(ctx *glib.CallbackContext) {
			if out != nil {
				*out = ctx.Data().(int8)
			}
			w.GetTopLevel().Destroy()
		}, def.Result)
		hbox.Add(b)
	}
	// TODO: right-align this one:
	b := gtk.NewButtonWithLabel(g.Right.Label)
	b.Clicked(func() {
		if out != nil {
			*out = g.Right.Result
		}
		w.GetTopLevel().Destroy()
	})
	hbox.Add(b)
	w.Add(hbox)
}

func (w GtkWindow) NewFrame(title string) Frame {
	var f GtkFrame
	f.Frame = gtk.NewFrame(title)
	f.VBox = gtk.NewVBox(true, 1) // (homogeneous bool, spacing int)
	f.Frame.Add(f.VBox)
	w.Add(f.Frame)
	return f
}

func (w GtkWindow) Show(width, height int) {
	w.Window.SetSizeRequest(width, height)
	w.Window.ShowAll()
	gtk.Main()
}
