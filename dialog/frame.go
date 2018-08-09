package dialog

type Frame interface {
	NewEntry(*string, string, bool)
	NewLabel(string)
}
