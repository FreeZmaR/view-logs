package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"log"
)

type logItemRenderer struct {
	objects    []fyne.CanvasObject
	background *canvas.Rectangle
	layout     fyne.Layout
	component  *LogItem
	level      fyne.CanvasObject
	msg        fyne.CanvasObject
	time       *LogTime
}

var _ fyne.WidgetRenderer = &logItemRenderer{}

func (r *logItemRenderer) Destroy() {
	log.Print("destroy")
}

func (r *logItemRenderer) Layout(size fyne.Size) {
	log.Print("layout: ", size)

	r.background.Resize(size)

	objects := []fyne.CanvasObject{r.level, r.msg, r.time}

	min := r.layout.MinSize(objects)
	r.layout.Layout(objects, min)

	sizeShift := size.Width - r.MinSize().Width

	position := fyne.Position{X: 0, Y: 0}
	levelPosition := r.level.Position().Add(fyne.Position{X: position.X + r.lvlPadding(), Y: position.Y})
	msgPosition := r.msg.Position().Add(fyne.Position{X: position.X + sizeShift/2})
	timePosition := r.time.Position().Add(fyne.Position{X: position.X + sizeShift})

	r.level.Move(levelPosition)
	r.msg.Move(msgPosition)
	r.time.Move(timePosition)
}

func (r *logItemRenderer) MinSize() fyne.Size {
	var (
		width          = r.lvlPadding()
		height float32 = 0
	)

	for _, object := range r.objects {
		size := object.MinSize()
		width += size.Width

		if height < size.Height {
			height = size.Height
		}
	}

	return fyne.NewSize(
		width,
		height,
	).Add(
		fyne.NewSize(
			theme.InnerPadding()*2,
			theme.InnerPadding()*2,
		),
	)
}

func (r *logItemRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *logItemRenderer) Refresh() {
	log.Print("refresh")

	r.background.Refresh()
	r.time.Refresh()
	r.Layout(r.component.Size())

	canvas.Refresh(r.component)
}

func (r *logItemRenderer) lvlPadding() float32 {
	return 20
}
