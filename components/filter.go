package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
)

type Filter struct {
	widget.BaseWidget
	items           []*FilterItem
	onSubmit        func(text string)
	itemsContainer  *fyne.Container
	window          fyne.Window
	dialog          *FilterDialog
	itemsController *fyne.Container
	checkboxAll     *widget.Check
}

func NewFilter(w fyne.Window, onSubmit func(text string), items ...*FilterItem) *Filter {
	f := &Filter{
		onSubmit: onSubmit,
		items:    items,
		window:   w,
	}
	f.makeFilterDialog()
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

	toolBar := widget.NewToolbar(
		widget.NewToolbarAction(
			theme.ContentAddIcon(),
			func() {
				comp.dialog.Show()
			},
		),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaReplayIcon(), func() {
			//TODO: history
		}),
	)

	comp.makeFilterItemsContainer()

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
}

func (comp *Filter) AddFilter(filterType FilterItemType, name string, value string) {
	item, err := NewFilterItem(
		filterType,
		name,
		value,
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

	comp.itemsController.Show()
	comp.AddItem(item)
}

func (comp *Filter) RemoveItem(position int) {
	item := comp.items[position]
	comp.itemsContainer.Remove(item)

	comp.items = append(comp.items[:position], comp.items[position+1:]...)

	for i := range comp.items {
		comp.items[i].position = i
	}

	if len(comp.items) == 0 {
		comp.checkboxAll.Checked = false
		comp.checkboxAll.Refresh()
		comp.itemsController.Hide()
	}

	comp.Refresh()
}

func (comp *Filter) RemoveSelected() {
	for i, item := range comp.items {
		if item.checkbox.Checked {
			comp.RemoveItem(i)
		}
	}

	comp.checkboxAll.Checked = false
	comp.checkboxAll.Refresh()
}

func (comp *Filter) DeactivateFilters() {
	for _, item := range comp.items {
		item.Deactivate()
	}
}

func (comp *Filter) ActivateFilters() {
	for _, item := range comp.items {
		item.Activate()
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

func (comp *Filter) makeFilterDialog() {
	comp.dialog = NewFilterDialog(comp.window, comp.AddFilter)
}

func (comp *Filter) makeFilterItemsContainer() {
	comp.checkboxAll = widget.NewCheck("All", func(b bool) {
		if !b {
			comp.DeactivateFilters()

			return
		}

		comp.ActivateFilters()
	})

	button := widget.NewButton("delete", func() {
		comp.RemoveSelected()
	})

	comp.itemsController = container.NewHBox(comp.checkboxAll, widget.NewSeparator(), button)
	comp.itemsController.Hide()

	comp.itemsContainer = container.NewVBox(comp.itemsController, widget.NewSeparator())
}
