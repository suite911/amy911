package widget

import (
	"github.com/g3n/engine/gui"
)

type Small struct {
	Panel *gui.Panel
	Label *gui.Label
}

func NewSmall(label string) (w *Small) {
	w = new(Small)
	w.Init(label)
	return
}

func (w *Small) Init(label string) {
	w.Panel = gui.NewPanel(0, 0)
	w.Label = gui.NewLabel(label)
	w.Panel.Add(w.Label)
	w.Panel.SetWidth(w.Label.TotalWidth())
	w.Panel.SetHeight(w.Label.TotalHeight())
}
