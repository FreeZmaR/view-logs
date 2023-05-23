package contents

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/FreeZmaR/view-logs/components"
	"github.com/FreeZmaR/view-logs/core/logger"
)

type Main struct {
	content    fyne.CanvasObject
	filter     *components.Filter
	mainWindow fyne.Window
}

func NewMain(w fyne.Window) *Main {
	main := Main{content: container.NewMax(), mainWindow: w}
	main.makeContent()

	return &main
}

func (c *Main) Content() fyne.CanvasObject {
	return c.content
}

func (c *Main) makeContent() {
	searchLabel := widget.NewLabel("Search")
	searchLabel.TextStyle.Bold = true
	searchLabel.Alignment = fyne.TextAlignLeading
	c.filter = components.NewFilter(
		c.mainWindow,
		func(text string) {
			fmt.Println("Filter text: " + text)
		},
	)

	header := container.NewVBox(
		container.NewVBox(
			searchLabel,
			widget.NewSeparator(),
		),
		c.filter,
	)

	logsLabel := widget.NewLabel("Logs")
	logs := container.NewBorder(
		container.NewVBox(
			logsLabel,
			widget.NewSeparator(),
		),
		nil,
		nil,
		nil,
		c.makeLogs(),
	)

	main := container.NewVBox(logs)
	content := container.NewVBox(header, widget.NewSeparator(), main)

	c.content = content
}

func (c *Main) makeLogs() fyne.CanvasObject {
	ms1 := `{"level": "info", "message": "First log", "time": "2021-08-15T15:04:05Z", "operation_id": "some_id"}`
	ms2 := `{"level": "warn", "message": "Second log", "time": "2021-08-15T15:05:05Z", "operation_id": "some_id"}`

	lg1 := logger.NewJSONLog()
	lg1.Parse([]byte(ms1))

	lg2 := logger.NewJSONLog()
	lg2.Parse([]byte(ms2))

	return container.NewVBox(
		components.NewLogItem(lg1),
		components.NewLogItem(lg2),
	)
}

func (c *Main) makeFilters() fyne.CanvasObject {
	return nil
}
