package widget

import (
	"strconv"

	"github.com/g3n/engine/gui"
)

type Performance struct {
	Inner, Outer *gui.Panel
	Units, Value *gui.Label
}

func NewPerformance(large int, label string) (w *Performance) {
	w = new(Performance)
	w.Init(large, label)
	return
}

func (w *Performance) Init(large int, label string) {
	w.Inner = gui.NewPanel(0, 0)
	w.Outer = gui.NewPanel(0, 0)
	w.Units = gui.NewLabel(label)
	w.Value = gui.NewLabel(strconv.Itoa(large))
	w.Outer.SetLayout(gui.NewDockLayout())
	w.Inner.SetLayoutParams(&gui.DockLayoutParams{gui.DockBottom})
	w.Outer.Add(w.Inner)
	w.Inner.Add(w.Units)
	w.Inner.Add(w.Value)
	uw, uh := float64(w.Units.TotalWidth()), float64(w.Units.TotalHeight())
	vw, vh := float64(w.Value.TotalWidth()), float64(w.Value.TotalHeight())
	width := float32(vw + uw)
	w.Inner.SetWidth(width)
	w.Outer.SetWidth(width)
	h := vh
	if uh > h {
		h = uh
	}
	w.Units.SetPosition(float32(vw), 0)
	height := float32(h)
	w.Inner.SetHeight(height)
	w.Outer.SetHeight(height)
	w.Value.SetText("")
}
