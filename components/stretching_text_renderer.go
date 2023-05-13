package components

import (
	"fyne.io/fyne/v2"
	fyneTheme "fyne.io/fyne/v2/theme"
)

type StretchingTextRenderer struct {
	component          *StretchingText
	text               fyne.CanvasObject
	minWidth, maxWidth float32
	objects            []fyne.CanvasObject
}

func newStretchingTextRenderer(
	component *StretchingText,
	text fyne.CanvasObject,
	minSize, maxSize int,
) *StretchingTextRenderer {
	return &StretchingTextRenderer{
		component: component,
		text:      text,
		minWidth:  float32(minSize) + fyneTheme.InnerPadding(),
		maxWidth:  float32(maxSize) + fyneTheme.InnerPadding(),
		objects:   []fyne.CanvasObject{text},
	}
}

func (r *StretchingTextRenderer) Destroy() {}

func (r *StretchingTextRenderer) Layout(size fyne.Size) {
	r.text.Resize(fyne.NewSize(size.Width, r.text.MinSize().Height))
	r.text.Move(fyne.NewPos(0, (size.Height-r.text.MinSize().Height)/2))
}

func (r *StretchingTextRenderer) MinSize() fyne.Size {
	return r.text.MinSize()
}

func (r *StretchingTextRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *StretchingTextRenderer) Refresh() {
	r.text.Refresh()
	r.Layout(r.component.Size())
}
