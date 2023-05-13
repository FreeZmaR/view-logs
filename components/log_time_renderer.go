package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type logTimeRenderer struct {
	component  *LogTime
	time, date fyne.CanvasObject
	layout     fyne.Layout
	objects    []fyne.CanvasObject
}

var _ fyne.WidgetRenderer = &logTimeRenderer{}

func (r *logTimeRenderer) Destroy() {}

func (r *logTimeRenderer) Layout(size fyne.Size) {
	objects := []fyne.CanvasObject{r.time, r.date}

	minSize := r.layout.MinSize(objects)
	r.layout.Layout(objects, minSize)

	objectsSize := r.MinSize()
	padding := r.padding()

	position := fyne.Position{
		X: (size.Width-objectsSize.Width)/2 + padding.Width/2,
		Y: (size.Height-objectsSize.Height)/2 + padding.Height/2,
	}

	timePosition := position.AddXY(0, -position.Y+padding.Height/2)
	datePosition := position.AddXY(0, -padding.Height/2)

	r.time.Move(r.time.Position().Add(timePosition))
	r.date.Move(r.date.Position().Add(datePosition))
}

func (r *logTimeRenderer) MinSize() fyne.Size {
	size := fyne.Size{}

	size.Width = fyne.Max(r.time.MinSize().Width, r.date.MinSize().Width)
	size.Height = r.time.MinSize().Height + r.date.MinSize().Height

	return size.Add(r.padding())
}

func (r *logTimeRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *logTimeRenderer) Refresh() {
	r.time.Refresh()
	r.date.Refresh()
	r.Layout(r.component.Size())

	canvas.Refresh(r.component)
}

// padding return fyne.Size for padding
func (r *logTimeRenderer) padding() fyne.Size {
	return fyne.NewSize(theme.InnerPadding()*2, theme.InnerPadding())
}
