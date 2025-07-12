package main

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
)

var w fyne.Window

func main() {
	a := fyneApp.NewWithID("4Go10")
	a.Settings().SetTheme(&nordTheme{})
	initStorage(a.Storage().RootURI())
	w = a.NewWindow("4Go10")
	w.SetContent(mainApp())
	w.ShowAndRun()
}
