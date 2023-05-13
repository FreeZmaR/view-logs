package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type FilterRenderer struct {
	component   *Filter
	searchInput *widget.Entry
	container   *fyne.Container
	objects     []fyne.CanvasObject
}

func (r *FilterRenderer) Destroy() {}

func (r *FilterRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *FilterRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *FilterRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *FilterRenderer) Refresh() {
	r.container.Refresh()
	canvas.Refresh(r.component)
}
