package dialog

var iLibrary ILibrary

func Library() ILibrary {
	return iLibrary
}

type ILibrary interface {
	NewWindow(string) Node
	NewEdit(Node, string, bool) Node
	NewFrame(Node, string) Node
	NewLabel(Node, string) Node
}
