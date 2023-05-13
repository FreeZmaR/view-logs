package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"strings"
	"time"
)

type Filter struct {
	widget.BaseWidget
	items          []*FilterItem
	onSubmit       func(text string)
	itemsContainer *fyne.Container
	window         fyne.Window
	counter        int
}

func NewFilter(w fyne.Window, onSubmit func(text string), items ...*FilterItem) *Filter {
	f := &Filter{
		onSubmit: onSubmit,
		items:    items,
		window:   w,
	}
	f.ExtendBaseWidget(f)

	return f
}

func (comp *Filter) CreateRenderer() fyne.WidgetRenderer {
	comp.ExtendBaseWidget(comp)

	searchInput := widget.NewEntry()
	searchInput.SetPlaceHolder("Any symbols from logs")
	searchInput.ActionItem = widget.NewIcon(theme.SearchIcon())
	searchInput.TextStyle.Bold = true
	searchInput.OnSubmitted = comp.onSubmit
	dialogWindow := comp.getDialogWindow()

	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.ContentAddIcon(),
			func() {
				dialogWindow.Show()
			},
		),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaReplayIcon(), func() {
			log.Print("Reset")
			comp.DeactivateFilters()
		}),
	)

	comp.itemsContainer = container.NewVBox()

	content := container.NewBorder(
		nil,
		comp.itemsContainer,
		nil,
		toolBar,
		searchInput,
	)

	return &FilterRenderer{
		component:   comp,
		searchInput: searchInput,
		container:   content,
		objects:     []fyne.CanvasObject{content},
	}
}

func (comp *Filter) AddItem(item *FilterItem) {
	item.position = len(comp.items)
	comp.items = append(comp.items, item)

	comp.itemsContainer.Add(item)
	comp.Refresh()

	fmt.Printf("Add item: pos: %d len: %d \n", item.position, len(comp.items))
}

func (comp *Filter) RemoveItem(position int) {
	item := comp.items[position]
	comp.itemsContainer.Remove(item)

	comp.items = append(comp.items[:position], comp.items[position+1:]...)

	for i := range comp.items {
		comp.items[i].position = i
	}

	comp.Refresh()
}

func (comp *Filter) DeactivateFilters() {
	for _, item := range comp.items {
		item.Deactivate()
	}
}

func (comp *Filter) HasKey(key string) bool {
	for _, item := range comp.items {
		if strings.ToLower(strings.TrimSpace(item.Key())) != strings.ToLower(strings.TrimSpace(key)) {
			return false
		}
	}

	return false
}

func (comp *Filter) IsEqual(key string, value any) bool {
	for _, item := range comp.items {
		if !item.IsActive() {
			continue
		}

		if strings.ToLower(strings.TrimSpace(item.Key())) != strings.ToLower(strings.TrimSpace(key)) {
			continue
		}

		if !item.IsEqual(value) {
			return false
		}
	}

	return true
}

func (comp *Filter) AddFilter(add bool) {
	if !add {
		return
	}

	item, err := NewFilterItem(
		FilterItemTypeLessThan,
		"operation_id_"+strconv.Itoa(comp.counter),
		time.Now().Format("15:04:05.000000"),
		comp.RemoveItem,
		func(position int) {
			log.Print("On active")

		},
		func(position int) {
			log.Print("On deactive")
		},
	)
	if err != nil {
		log.Print("Error: " + err.Error())

		return
	}

	comp.counter++

	comp.AddItem(item)
}

func (comp *Filter) getDialogWindow() dialog.Dialog {
	name := &widget.FormItem{
		Text:     "Field Name",
		Widget:   widget.NewEntry(),
		HintText: "Filed for search by filter",
	}

	filterType := &widget.FormItem{
		Text:     "Filter type",
		Widget:   widget.NewSelect(getFilterItemTypeText(), func(s string) {}),
		HintText: "Type of filter to search by rules",
	}

	value := &widget.FormItem{
		Text:     "Default value",
		Widget:   widget.NewEntry(),
		HintText: "Start value for filter",
	}

	return dialog.NewForm(
		"Add filter",
		"Add",
		"Cancel",
		[]*widget.FormItem{name, filterType, value},
		comp.AddFilter,
		comp.window,
	)
}
