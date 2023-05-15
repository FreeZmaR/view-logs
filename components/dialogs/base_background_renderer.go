package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type BaseBackgroundRenderer struct {
	rect    *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (r *BaseBackgroundRenderer) Destroy() {}

func (r *BaseBackgroundRenderer) Layout(size fyne.Size) {
	r.rect.Resize(size)
}

func (r *BaseBackgroundRenderer) MinSize() fyne.Size {
	return r.rect.MinSize()
}

func (r *BaseBackgroundRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *BaseBackgroundRenderer) Refresh() {
	red, g, b, _ := theme.OverlayBackgroundColor().RGBA()
	bg := &color.NRGBA{R: uint8(red), G: uint8(g), B: uint8(b), A: 230}
	r.rect.FillColor = bg
}
