package dialog

type Frame interface {
	NewButtonClose(*int, string)
	NewButtonGroup(*int, string)
	NewEntry(*string, string, bool)
	NewLabel(string)
}
