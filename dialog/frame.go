package dialog

type Frame interface {
	NewButtonGroup(*int8, *ButtonGroup)
	NewEntry(*string, bool)
	NewLabel(string)
}
