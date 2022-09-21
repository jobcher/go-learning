package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("激活软件")

	hello := widget.NewLabel("hello")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Action", func() {
			hello.SetText("this is fault")
		}),
		widget.NewButton("Return", func() {
			hello.SetText("hello")
		}),
	))

	w.ShowAndRun()
}
