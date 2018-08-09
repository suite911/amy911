package dialog

import (
	"os"

	"github.com/mattn/go-gtk"
)

func init() {
	Init = func() {
		gtk.Init(&os.Args)
		Library = LibraryGtk{}
	}
	Quit = func() {
	}
	Init()
}

type LibraryGtk struct {
}
