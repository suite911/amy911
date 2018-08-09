package dialog

type Frame interface {
	NewButtonClose(string)
	NewButtonGroup(*int, string)
	NewEntry(*string, string, bool)
	NewLabel(string)
}
