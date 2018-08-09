package dialog

type Frame interface {
	NewButtonGroup(*int, *ButtonGroup)
	NewEntry(*string, string, bool)
	NewLabel(string)
}
