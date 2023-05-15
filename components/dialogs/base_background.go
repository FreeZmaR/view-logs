package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type BaseBackground struct {
	widget.BaseWidget
}

func NewBaseBackground() *BaseBackground {
	bb := &BaseBackground{}
	bb.ExtendBaseWidget(bb)

	return bb
}

func (bb *BaseBackground) CreateRenderer() fyne.WidgetRenderer {
	bb.ExtendBaseWidget(bb)
	rect := canvas.NewRectangle(theme.OverlayBackgroundColor())

	return &BaseBackgroundRenderer{rect, []fyne.CanvasObject{rect}}
}
