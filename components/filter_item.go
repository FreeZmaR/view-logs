package components

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"regexp"
	"strconv"
	"strings"
)

const (
	FilterItemTypeEqual FilterItemType = iota + 1
	FilterItemTypeNotEqual
	FilterItemTypeContain
	FilterItemTypeNotContain
	FilterItemTypeRegexp
	FilterItemTypeGreatThan
	FilterItemTypeLessThan
)

type FilterItemType int

type FilterItem struct {
	widget.BaseWidget
	itemType     FilterItemType
	key          string
	position     int
	value        any
	isActive     bool
	regexp       *regexp.Regexp
	checkbox     *widget.Check
	onDelete     func(position int)
	onActive     func(position int)
	onDeactivate func(position int)
}

func NewFilterItem(
	itemType FilterItemType,
	key string,
	value any,
	onDelete func(position int),
	onActive func(position int),
	onDeactivate func(position int),
) (*FilterItem, error) {
	item := &FilterItem{
		itemType:     itemType,
		key:          key,
		value:        value,
		isActive:     false,
		onDelete:     onDelete,
		onActive:     onActive,
		onDeactivate: onDeactivate,
	}
	item.ExtendBaseWidget(item)

	if itemType == FilterItemTypeRegexp {
		v, ok := value.(string)
		if !ok {
			return nil, errors.New("regexp type: value must be string")
		}

		item.regexp = regexp.MustCompile(v)
	}

	return item, nil
}

func (comp *FilterItem) CreateRenderer() fyne.WidgetRenderer {
	comp.ExtendBaseWidget(comp)

	comp.checkbox = widget.NewCheck("", func(value bool) {
		if value {
			comp.Activate()
		} else {
			comp.Deactivate()
		}
	})

	label := widget.NewLabel(comp.key)
	label.TextStyle.Bold = true

	var input fyne.CanvasObject

	switch t := comp.value.(type) {
	case string:
		tmp := widget.NewEntry()
		tmp.Text = t
		input = tmp
	case int:
		tmp := widget.NewEntry()
		tmp.Text = strconv.Itoa(t)
		input = tmp
	case float32:
		tmp := widget.NewEntry()
		tmp.Text = strconv.FormatFloat(float64(t), 'f', -1, 32)
		input = tmp
	case bool:
		tmp := widget.NewEntry()

		if t {
			tmp.Text = "true"
		} else {
			tmp.Text = "false"
		}

		input = tmp
	}

	marker := comp.getFilterItemMarker()

	button := widget.NewButtonWithIcon("", theme.DeleteIcon(), comp.OnDelete)

	return &FilterItemRenderer{
		component:  comp,
		layout:     layout.NewHBoxLayout(),
		checkbox:   comp.checkbox,
		input:      input,
		label:      label,
		button:     button,
		marker:     marker,
		objects:    []fyne.CanvasObject{comp.checkbox, label, marker, input, button},
		inPosition: false,
	}
}

func (comp *FilterItem) Activate() {
	comp.isActive = true
	comp.checkbox.Checked = true
	comp.OnActive()
	comp.checkbox.Refresh()
}

func (comp *FilterItem) Deactivate() {
	comp.isActive = false
	comp.checkbox.Checked = false
	comp.OnDeactivate()
	comp.checkbox.Refresh()
}

func (comp *FilterItem) OnDelete() {
	if comp.onDelete != nil {
		comp.onDelete(comp.position)
	}
}

func (comp *FilterItem) OnActive() {
	if comp.onActive != nil {
		comp.onActive(comp.position)
	}
}

func (comp *FilterItem) OnDeactivate() {
	if comp.onDeactivate != nil {
		comp.onDeactivate(comp.position)
	}
}

func (comp *FilterItem) Key() string {
	return comp.key
}

func (comp *FilterItem) SetPosition(position int) {
	comp.position = position
}

func (comp *FilterItem) IsActive() bool {
	return comp.isActive
}

func (comp *FilterItem) IsEqual(val any) bool {
	switch comp.itemType {
	case FilterItemTypeEqual:
		return comp.value == val
	case FilterItemTypeNotEqual:
		return comp.value != val
	case FilterItemTypeContain:
		v1, ok1 := val.(string)
		v2, ok2 := comp.value.(string)
		if !ok1 || !ok2 {
			return false
		}

		return strings.Contains(v1, v2)
	case FilterItemTypeNotContain:
		v1, ok1 := val.(string)
		v2, ok2 := comp.value.(string)
		if !ok1 || !ok2 {
			return true
		}

		return !strings.Contains(v1, v2)
	case FilterItemTypeRegexp:
		v1, ok1 := val.(string)
		if !ok1 || comp.regexp == nil {
			return false
		}

		return comp.regexp.MatchString(v1)
	case FilterItemTypeGreatThan:
		switch t := val.(type) {
		case int:
			v, ok := comp.value.(int)
			if !ok {
				return false
			}

			return t > v
		case float32:
			v, ok := comp.value.(float32)
			if !ok {
				return false
			}

			return t > v
		default:
			return false
		}
	case FilterItemTypeLessThan:
		switch t := val.(type) {
		case int:
			v, ok := comp.value.(int)
			if !ok {
				return false
			}

			return t < v
		case float32:
			v, ok := comp.value.(float32)
			if !ok {
				return false
			}

			return t < v
		default:
			return false
		}
	default:
		return false
	}
}

func (comp *FilterItem) getFilterItemMarker() fyne.CanvasObject {
	var text string

	switch comp.itemType {
	case FilterItemTypeEqual:
		text = "="
	case FilterItemTypeNotEqual:
		text = "!="
	case FilterItemTypeContain:
		text = "contain"
	case FilterItemTypeNotContain:
		text = "not contain"
	case FilterItemTypeRegexp:
		text = "regexp"
	case FilterItemTypeGreatThan:
		text = ">"
	case FilterItemTypeLessThan:
		text = "<"
	default:
	}

	marker := canvas.NewText(text, color.RGBA{R: 114, G: 17, B: 17, A: 255})
	marker.TextStyle.Bold = true
	marker.TextStyle.Italic = true
	marker.TextSize = 18

	return marker
}

func getFilterItemTypeText() []string {
	return []string{
		"equal",
		"not equal",
		"contain",
		"not contain",
		"regexp",
		"greater than",
		"less than",
	}
}
