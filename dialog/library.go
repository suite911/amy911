package dialog

var iLibrary ILibrary

func Library() ILibrary {
	return iLibrary
}

type ILibrary interface {
	NewWindow(string) Window
}
