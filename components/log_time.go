package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	fyneWidget "fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

// LogTime is a widget that displays a log item's time.
type LogTime struct {
	fyneWidget.BaseWidget
	time string
	date string
}

// NewLogTime creates a new LogTime widget.
func NewLogTime(t time.Time) *LogTime {
	l := &LogTime{
		time: t.Format("15:04:05"),
		date: t.Format("2006-01-02"),
	}

	l.ExtendBaseWidget(l)

	return l
}

// CreateRenderer implements fyne.Widget.
func (comp *LogTime) CreateRenderer() fyne.WidgetRenderer {
	comp.ExtendBaseWidget(comp)

	timeWidget := comp.getTimeWidget()
	dateWidget := comp.getDateWidget()

	return &logTimeRenderer{
		component: comp,
		time:      timeWidget,
		date:      dateWidget,
		layout:    layout.NewVBoxLayout(),
		objects: []fyne.CanvasObject{
			timeWidget,
			dateWidget,
		},
	}
}

// getTimeWidget returns the fyne.CanvasObject for the time.
func (comp *LogTime) getTimeWidget() fyne.CanvasObject {
	text := canvas.NewText(comp.time, comp.textColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true
	text.TextStyle.Monospace = true
	text.TextSize = 14

	return text
}

// getDateWidget returns the fyne.CanvasObject for the date.
func (comp *LogTime) getDateWidget() fyne.CanvasObject {
	text := canvas.NewText(comp.date, comp.textColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true
	text.TextStyle.Monospace = true
	text.TextSize = 10

	return text
}

// textColor returns the color of the text.
func (comp *LogTime) textColor() color.Color {
	return color.RGBA{R: 180, G: 180, B: 180, A: 255}
}
