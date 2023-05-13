package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	fyneTheme "fyne.io/fyne/v2/theme"
	fyneWidget "fyne.io/fyne/v2/widget"
	"strings"
	"time"
)

const (
	LogItemLevelUnknown LogItemLevel = 0
	LogItemLevelInfo    LogItemLevel = 1
	LogItemLevelError   LogItemLevel = 2
	LogItemLevelWarning LogItemLevel = 3
	LogItemLevelFatal   LogItemLevel = 4
)

type LogItemI interface {
	fyne.Widget
	desktop.Hoverable
}

type LogItem struct {
	fyneWidget.BaseWidget
	msg            string
	level          LogItemLevel
	time           time.Time
	fieldContainer *logItemFieldContainer
	background     *canvas.Rectangle
}

type LogItemLevel int

var _ LogItemI = (*LogItem)(nil)

func NewLogItem(msg, level string, t time.Time, fields ...LogItemField) *LogItem {
	l := &LogItem{
		msg:            msg,
		level:          parseItemLogLevel(level),
		time:           t,
		fieldContainer: makeLogItemFieldContainer(fields),
	}

	l.ExtendBaseWidget(l)

	return l
}

func (comp *LogItem) CreateRenderer() fyne.WidgetRenderer {
	comp.ExtendBaseWidget(comp)

	lvlWidget := comp.getLevelWidget()
	msgWidget := NewStretchingText(comp.msg, 15, len(comp.msg))
	timeWidget := NewLogTime(comp.time)

	comp.background = canvas.NewRectangle(fyneTheme.ErrorColor())
	comp.background.FillColor = fyneTheme.BackgroundColor()

	r := &logItemRenderer{
		objects: []fyne.CanvasObject{
			comp.background,
			lvlWidget,
			msgWidget,
			timeWidget,
		},
		background: comp.background,
		layout:     layout.NewHBoxLayout(),
		component:  comp,
		level:      lvlWidget,
		msg:        msgWidget,
		time:       timeWidget,
	}

	return r
}

func (comp *LogItem) MouseIn(*desktop.MouseEvent) {
	comp.background.FillColor = fyneTheme.HoverColor()
	comp.background.Refresh()
}

func (comp *LogItem) MouseMoved(event *desktop.MouseEvent) {
}

func (comp *LogItem) MouseOut() {
	comp.background.FillColor = fyneTheme.BackgroundColor()
	comp.background.Refresh()
}

func (comp *LogItem) getLevelWidget() fyne.CanvasObject {
	var (
		text  string
		color string
	)

	switch comp.level {
	case LogItemLevelInfo:
		text = "[ Info ]"
		color = fyneTheme.ColorBlue
	case LogItemLevelError:
		text = "[ Error ]"
		color = fyneTheme.ColorRed
	case LogItemLevelWarning:
		text = "[ Warning ]"
		color = fyneTheme.ColorOrange
	case LogItemLevelFatal:
		text = "[ Fatal ]"
		color = fyneTheme.ColorPurple
	case LogItemLevelUnknown:
		text = "[ Unknown ]"
		color = fyneTheme.ColorGray
	}

	canvasText := canvas.NewText(text, fyneTheme.PrimaryColorNamed(color))
	canvasText.Alignment = fyne.TextAlignCenter
	canvasText.TextStyle.Bold = true
	canvasText.TextStyle.Italic = true

	return canvasText
}

func parseItemLogLevel(rawLevel string) LogItemLevel {
	level := strings.ToLower(rawLevel)

	switch level {
	case "info":
		return LogItemLevelInfo
	case "error":
		return LogItemLevelError
	case "warning":
		return LogItemLevelWarning
	case "fatal":
		return LogItemLevelFatal
	default:
		return LogItemLevelUnknown
	}
}

type logItemFieldContainer struct {
	fields []LogItemField
}

func makeLogItemFieldContainer(fields []LogItemField) *logItemFieldContainer {
	f := &logItemFieldContainer{fields: make([]LogItemField, 0, len(fields))}
	f.addUniqueFields(fields)

	return f
}

func (container *logItemFieldContainer) addUniqueFields(fields []LogItemField) {
	uniqueNewFields := make([]LogItemField, 0, len(fields))

	for _, field := range fields {
		isUnique := true

		for _, existField := range container.fields {
			if existField.Label == field.Label {
				isUnique = false

				break
			}
		}

		if isUnique {
			uniqueNewFields = append(uniqueNewFields, field)
		}
	}

	container.fields = append(container.fields, uniqueNewFields...)
}

func (container *logItemFieldContainer) showByLabel(label string) {
	container.changeIsShowByLabel(label, true)
}

func (container *logItemFieldContainer) hideByLabel(label string) {
	container.changeIsShowByLabel(label, false)
}

func (container *logItemFieldContainer) changeIsShowByLabel(label string, isShow bool) {
	for _, field := range container.fields {
		if field.Label == label {
			field.isShow = isShow
		}
	}
}

type LogItemField struct {
	isShow bool
	Label  string
	Value  string
}

func NewLogItemField(label, value string) LogItemField {
	return LogItemField{Label: label, Value: value, isShow: false}
}
