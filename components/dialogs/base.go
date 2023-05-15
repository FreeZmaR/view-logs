package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Base struct {
	title       string
	icon        fyne.Resource
	desiredSize fyne.Size

	window         *widget.PopUp
	background     *BaseBackground
	content, label fyne.CanvasObject
	dismiss        *widget.Button
	parent         fyne.Window
	layout         *Layout
}

var _ dialog.Dialog = (*Base)(nil)

func NewBase(title, dismiss string, content fyne.CanvasObject, parent fyne.Window) *Base {
	instance := &Base{
		content:    content,
		title:      title,
		icon:       nil,
		parent:     parent,
		background: NewBaseBackground(),
		label:      widget.NewLabelWithStyle(title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	}
	instance.layout = &Layout{instance: instance}

	instance.dismiss = &widget.Button{Text: dismiss,
		OnTapped: instance.Hide,
	}

	modalContent := container.New(
		instance.layout,
		&canvas.Image{Resource: instance.icon},
		instance.background,
		instance.content,
		container.NewHBox(layout.NewSpacer(), instance.dismiss, layout.NewSpacer()),
		instance.label,
	)
	instance.window = widget.NewModalPopUp(modalContent, parent.Canvas())

	return instance
}

func (comp *Base) SetDismissText(label string) {
	comp.dismiss.Text = label
	comp.window.Refresh()
}

func (comp *Base) SetOnClosed(closed func()) {

}

func (comp *Base) MinSize() fyne.Size {
	return comp.window.MinSize()
}

func (comp *Base) Resize(size fyne.Size) {
	comp.desiredSize = size
	comp.window.Resize(size)
}

func (comp *Base) Show() {
	comp.window.Show()
}

func (comp *Base) Hide() {
	comp.window.Hide()
}

func (comp *Base) Refresh() {
	comp.window.Refresh()
}
