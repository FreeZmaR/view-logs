package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/FreeZmaR/view-logs/components/dialogs"
	"github.com/FreeZmaR/view-logs/components/forms"
	"github.com/FreeZmaR/view-logs/core/logger"
)

type FilterDialog struct {
	instance   *dialogs.Form
	window     fyne.Window
	name       *widget.Entry
	filterType *widget.Select
	value      *widget.Entry
	onSubmit   FilterDialogSubmit
}

type FilterDialogSubmit func(filterType logger.CompareType, name string, value string)

func NewFilterDialog(window fyne.Window, onSubmit FilterDialogSubmit) *FilterDialog {
	filterDialog := &FilterDialog{window: window, onSubmit: onSubmit}

	return filterDialog.make()
}

func (comp *FilterDialog) Submit() {
	comp.onSubmit(logger.GetCompareTypeByString(comp.filterType.Selected), comp.name.Text, comp.value.Text)
	comp.Reset()
}

func (comp *FilterDialog) Reset() {
	comp.filterType.Selected = ""
	comp.name.Text = ""
	comp.value.Text = ""
}

func (comp *FilterDialog) Show() {
	comp.instance.Show()
}

func (comp *FilterDialog) Hide() {
	comp.instance.Hide()
}

func (comp *FilterDialog) make() *FilterDialog {
	comp.instance = dialogs.NewForm(
		"Add filter",
		"Add",
		"Cancel",
		comp.makeContent(),
		comp.window,
		comp.Submit,
		comp.Reset,
	)
	comp.instance.SetValidateBefore(true)
	comp.activateValidation()

	return comp
}

func (comp *FilterDialog) makeContent() []*forms.FormItem {
	comp.name = widget.NewEntry()
	comp.value = widget.NewEntry()
	comp.filterType = widget.NewSelect(logger.GetCompareTypeText(), func(s string) {})

	return []*forms.FormItem{
		{
			FormItem: &widget.FormItem{
				Text:     "Field Name",
				Widget:   comp.name,
				HintText: "Filed for search by filter",
			},
			Validation: func(item *forms.FormItem) bool {
				return item.Widget.(*widget.Entry).Text != ""
			},
		},
		{
			FormItem: &widget.FormItem{
				Text:     "Filter type",
				Widget:   comp.filterType,
				HintText: "Type of filter to search by rules",
			},
			Validation: func(item *forms.FormItem) bool {
				return item.Widget.(*widget.Select).Selected != ""
			},
		},
		{
			FormItem: &widget.FormItem{
				Text:     "Default value",
				Widget:   comp.value,
				HintText: "Start value for filter",
			},
		},
	}
}

func (comp *FilterDialog) activateValidation() {
	comp.name.OnChanged = func(s string) {
		comp.instance.Validate()
	}
	comp.filterType.OnChanged = func(s string) {
		comp.instance.Validate()
	}
}
