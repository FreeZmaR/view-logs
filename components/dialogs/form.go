package dialogs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/FreeZmaR/view-logs/components/forms"
)

type Form struct {
	*Base
	onSubmit         func()
	onCanceled       func()
	items            []*forms.FormItem
	ConfirmBtn       *widget.Button
	CancelBtn        *widget.Button
	closeWithContent bool
	validateBefore   bool
}

func NewForm(title, confirm, cancel string, items []*forms.FormItem, parent fyne.Window, onConfirm, onCancel func()) *Form {
	form := widget.NewForm()
	for _, item := range items {
		form.Append(item.FormItem.Text, item.FormItem.Widget)
	}

	instance := &Form{
		Base:             NewBase(title, cancel, form, parent),
		items:            items,
		closeWithContent: true,
		validateBefore:   false,
		onSubmit:         onConfirm,
		onCanceled:       onCancel,
	}
	instance.Base.layout = &Layout{instance: instance.Base}

	instance.ConfirmBtn = &widget.Button{Text: confirm, OnTapped: instance.Submit, Importance: widget.HighImportance}
	instance.CancelBtn = &widget.Button{Text: cancel, OnTapped: instance.Cancel}

	modalContent := container.New(
		instance.layout,
		&canvas.Image{Resource: instance.icon},
		instance.background,
		instance.content,
		container.NewHBox(layout.NewSpacer(), instance.ConfirmBtn, instance.CancelBtn, layout.NewSpacer()),
		instance.label,
	)
	instance.window = widget.NewModalPopUp(modalContent, parent.Canvas())

	return instance
}

func (comp *Form) SetValidateBefore(validate bool) {
	comp.validateBefore = validate

	if comp.validateBefore {
		comp.ConfirmBtn.Disable()
	}
}

func (comp *Form) Submit() {
	if nil == comp.onSubmit || comp.ConfirmBtn.Disabled() {
		return
	}

	comp.onSubmit()
	comp.Hide()
	if comp.validateBefore {
		comp.ConfirmBtn.Disable()
	}
}

func (comp *Form) Cancel() {
	if comp.onCanceled != nil {
		comp.onCanceled()
	}

	comp.Hide()

	if comp.validateBefore {
		comp.ConfirmBtn.Disable()
	}
}

func (comp *Form) Validate() {
	for _, item := range comp.items {
		if item.Validation != nil && !item.Validation(item) {
			comp.ConfirmBtn.Disable()

			return
		}
	}

	comp.ConfirmBtn.Enable()
}
