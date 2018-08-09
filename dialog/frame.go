package dialog

type Frame interface {
	NewButtonRow(*bool, bool)
	NewEntry(*string, string, bool)
	NewLabel(string)
}
