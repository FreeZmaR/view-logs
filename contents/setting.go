package contents

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Setting struct {
	content *fyne.Container
}

func NewSetting() *Setting {
	setting := Setting{content: container.NewMax()}

	return &setting
}

func (c *Setting) Content() *fyne.Container {
	return c.content
}
