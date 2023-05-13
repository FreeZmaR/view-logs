package contents

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/FreeZmaR/view-logs/components"
	"time"
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
	return container.NewVBox(
		components.NewLogItem(
			"First log",
			"info",
			time.Now().Add(-5*time.Minute),
			components.LogItemField{
				Label: "operation_id",
				Value: "some_id",
			},
		),
		components.NewLogItem(
			"Second log",
			"info",
			time.Now(),
			components.LogItemField{
				Label: "operation_id",
				Value: "some_id",
			},
		),
	)
}

func (c *Main) makeFilters() fyne.CanvasObject {
	return nil
}
