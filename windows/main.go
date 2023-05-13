package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/FreeZmaR/view-logs/contents"
	"log"
)

type Main struct {
	app      fyne.App
	window   fyne.Window
	menu     *fyne.MainMenu
	contents *mainContent
}

type mainContent struct {
	main    *contents.Main
	setting *contents.Setting
}

func NewMain(app fyne.App) *Main {
	window := app.NewWindow("Main")
	main := Main{
		app:    app,
		window: window,
		contents: &mainContent{
			main:    contents.NewMain(window),
			setting: contents.NewSetting(),
		},
	}

	main.makeTray()
	main.setMenu()
	main.window.SetMaster()
	main.window.SetIcon(theme.FyneLogo())

	return &main
}

func (m *Main) ShowAndRun() {
	m.window.SetContent(m.contents.main.Content())
	m.window.Resize(fyne.NewSize(640, 460))
	m.window.CenterOnScreen()
	m.window.ShowAndRun()
}

func (m *Main) makeTray() {
	if desk, ok := m.app.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}

		desk.SetSystemTrayMenu(menu)
	}
}

func (m *Main) setMenu() {
	settings := fyne.NewMenu("Settings", fyne.NewMenuItem("Auth", func() {
		log.Println("Setting.Auth")
	}))

	help := fyne.NewMenu("Help", fyne.NewMenuItem("About", func() {
		log.Println("Help.About")
	}))

	m.menu = fyne.NewMainMenu(
		settings,
		help,
	)

	m.window.SetMainMenu(m.menu)
}
