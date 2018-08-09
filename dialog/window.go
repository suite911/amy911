package dialog

type Window interface {
	NewButtonGroup(*int8, *ButtonGroup)
	NewFrame(string) Frame
	Show(int, int)
}
