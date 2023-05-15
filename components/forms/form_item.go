package forms

import "fyne.io/fyne/v2/widget"

type FormItemType interface {
	*widget.Entry | *widget.Select | *widget.Check
}

type FormItem struct {
	*widget.FormItem
	Validation func(item *FormItem) bool
}
