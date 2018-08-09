package dialog

import (
	"os"

	"github.com/mattn/go-gtk"
)

func init() {
	SetUp(
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
}
