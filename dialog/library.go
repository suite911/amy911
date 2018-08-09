package dialog

var iLibrary ILibrary

func Library() ILibrary {
	return iLibrary
}

type ILibrary interface {
	NewEntry(Node, string, bool)
	NewFrame(Node, string) Node
	NewLabel(Node, string)
	NewWindow(string) Node
}
