package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type FilterItemRenderer struct {
	component                              *FilterItem
	layout                                 fyne.Layout
	checkbox, input, label, button, marker fyne.CanvasObject
	objects                                []fyne.CanvasObject
	inPosition                             bool
}

func (r *FilterItemRenderer) Destroy() {}

func (r *FilterItemRenderer) Layout(size fyne.Size) {
	minSize := r.layout.MinSize(r.objects)
	r.layout.Layout(r.objects, minSize)

	if r.input.Size().Width < 150 {
		r.input.Resize(fyne.Size{Width: 250, Height: r.input.Size().Height})
		r.input.Refresh()
	}

	position := fyne.Position{}

	labelPosition := r.label.Position().Add(fyne.NewPos(position.X-theme.InnerPadding()*2, 0))
	markerPosition := r.marker.Position().Add(fyne.NewPos(position.X-theme.InnerPadding()*3, 0))
	inputPosition := r.input.Position().Add(fyne.NewPos(position.X-theme.InnerPadding()*2, 0))
	buttonPosition := r.button.Position().Add(fyne.NewPos(r.input.Size().Width-theme.InnerPadding()*6, 0))

	r.checkbox.Move(r.checkbox.Position().Add(position))
	r.label.Move(labelPosition)
	r.marker.Move(markerPosition)
	r.input.Move(inputPosition)
	r.button.Move(buttonPosition)
}

func (r *FilterItemRenderer) MinSize() fyne.Size {
	size := fyne.Size{}
	var height float32

	for _, object := range r.objects {
		size.Width = object.MinSize().Width
		height = fyne.Max(height, object.MinSize().Height)
	}

	if r.input.MinSize().Width < 150 {
		size.Width += 150 - r.input.MinSize().Width
	}

	size.Height = height

	return size.Add(fyne.Size{Width: theme.InnerPadding() * 2, Height: theme.InnerPadding()})
}

func (r *FilterItemRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *FilterItemRenderer) Refresh() {
	r.checkbox.Refresh()
	r.input.Refresh()
	r.label.Refresh()
	r.button.Refresh()
	canvas.Refresh(r.component)
}
