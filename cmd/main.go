package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/FreeZmaR/view-logs/windows"
)

func main() {
	a := app.New()
	window := windows.NewMain(a)
	window.ShowAndRun()
}

//func main() {
//	myApp := app.New()
//	myWindow := myApp.NewWindow("Border Layout")
//
//	top := canvas.NewText("top bar", color.White)
//	left := canvas.NewText("left", color.White)
//	rigth := canvas.NewText("right", color.White)
//	bottom := canvas.NewText("bottom", color.White)
//	middle := canvas.NewText("content", color.White)
//	content := container.NewBorder(top, bottom, left, rigth, middle)
//	myWindow.SetContent(content)
//	myWindow.ShowAndRun()
//}
