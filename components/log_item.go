package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	fyneTheme "fyne.io/fyne/v2/theme"
	fyneWidget "fyne.io/fyne/v2/widget"
	"github.com/FreeZmaR/view-logs/core/logger"
	"time"
)

type LogItemI interface {
	fyne.Widget
	desktop.Hoverable
}

type LogItem struct {
	fyneWidget.BaseWidget
	msg        string
	level      logger.Level
	time       time.Time
	background *canvas.Rectangle
	log        logger.Log
}

type LogItemLevel int

var _ LogItemI = (*LogItem)(nil)

func NewLogItem(log logger.Log) *LogItem {
	l := &LogItem{
		log:   log,
		level: log.GetLevel(),
		msg:   log.GetMessage(),
		time:  log.GetTime(),
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
	case logger.LevelInfo:
		text = "[ Info ]"
		color = fyneTheme.ColorBlue
	case logger.LevelError:
		text = "[ Error ]"
		color = fyneTheme.ColorRed
	case logger.LevelWarning:
		text = "[ Warning ]"
		color = fyneTheme.ColorOrange
	case logger.LevelFatal:
		text = "[ Fatal ]"
		color = fyneTheme.ColorPurple
	case logger.LevelUnknown:
		text = "[ Unknown ]"
		color = fyneTheme.ColorGray
	}

	canvasText := canvas.NewText(text, fyneTheme.PrimaryColorNamed(color))
	canvasText.Alignment = fyne.TextAlignCenter
	canvasText.TextStyle.Bold = true
	canvasText.TextStyle.Italic = true

	return canvasText
}
