package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("4Go10")

	w.SetContent(widget.NewLabel("Hi"))
	w.ShowAndRun()
}
