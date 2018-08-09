package dialog

type Frame interface {
	NewEntry(*string, bool)
	NewLabel(string)
}
