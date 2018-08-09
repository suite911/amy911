package dialog

type Window interface {
	NewFrame(string) Frame
	Show(int, int)
}
