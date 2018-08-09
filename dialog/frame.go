package dialog

type Frame interface {
	NewButtonRow(*bool, string)
	NewEntry(*string, string, bool)
	NewLabel(string)
}
