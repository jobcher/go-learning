package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {

	var (
		myApp    fyne.App
		myWindow fyne.Window
	)

	myApp = app.New()
	myWindow = myApp.NewWindow("hello")
	myWindow.SetContent(widget.NewLabel("hello"))

	myWindow.Show()
	myApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Print("exited")
}
