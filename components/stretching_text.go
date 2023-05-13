package components

import (
	"fyne.io/fyne/v2"
	fyneTheme "fyne.io/fyne/v2/theme"
	fyneWidget "fyne.io/fyne/v2/widget"
)

type StretchingText struct {
	fyneWidget.BaseWidget
	text             string
	minSize, maxSize int
}

func NewStretchingText(text string, minSize, maxSize int) *StretchingText {
	s := &StretchingText{
		text:    text,
		minSize: minSize,
		maxSize: maxSize,
	}

	s.ExtendBaseWidget(s)

	return s
}

func (comp *StretchingText) CreateRenderer() fyne.WidgetRenderer {
	comp.ExtendBaseWidget(comp)

	return newStretchingTextRenderer(comp, comp.getTextWidget(), comp.minSize, comp.maxSize)
}

func (comp *StretchingText) getTextWidget() fyne.CanvasObject {
	seg := &fyneWidget.TextSegment{Text: comp.text, Style: fyneWidget.RichTextStyleStrong}
	seg.Style.Alignment = fyne.TextAlignCenter
	seg.Style.Inline = true

	widget := fyneWidget.NewRichText(seg)
	widget.Resize(fyne.NewSize(fyneTheme.InnerPadding(), fyneTheme.InnerPadding()))

	return widget
}
