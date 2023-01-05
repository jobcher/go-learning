package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var (
		a     fyne.App
		w     fyne.Window
		hello *widget.Label
	)

	a = app.New()
	w = a.NewWindow("Hello")

	hello = widget.NewLabel("Hello jobcher")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
		widget.NewButton("return", func() {
			hello.SetText("Hello jobcher")
		}),
	))

	w.ShowAndRun()
}
