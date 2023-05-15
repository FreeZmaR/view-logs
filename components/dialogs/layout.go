package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Layout struct {
	instance *Base
}

var _ fyne.Layout = (*Layout)(nil)

func (l *Layout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	l.instance.background.Move(fyne.NewPos(0, 0))
	l.instance.background.Resize(size)

	buttons := objects[3]
	buttonsMin := buttons.MinSize()
	buttons.Resize(buttonsMin)
	buttons.Move(fyne.NewPos(
		size.Width/2-(buttonsMin.Width/2),
		size.Height-16-buttonsMin.Height,
	))

	contentStart := l.instance.label.Position().Y + l.instance.label.MinSize().Height + 16
	contentEnd := buttons.Position().Y - theme.Padding()
	objects[2].Move(fyne.NewPos(16/2, l.instance.label.MinSize().Height+16))
	objects[2].Resize(fyne.NewSize(size.Width-16, contentEnd-contentStart))
}

func (l *Layout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	contentMin := objects[2].MinSize()
	btnMin := objects[3].MinSize()

	width := fyne.Max(fyne.Max(contentMin.Width, btnMin.Width), objects[4].MinSize().Width) + 16
	height := contentMin.Height + btnMin.Height + l.instance.label.MinSize().Height + theme.Padding() + 16*2

	return fyne.NewSize(width, height)
}
