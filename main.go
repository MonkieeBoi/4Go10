package main

import (
	fyneApp "fyne.io/fyne/v2/app"
)

func main() {
	a := fyneApp.New()
	a.Settings().SetTheme(&nordTheme{})
	w := a.NewWindow("4Go10")
	w.SetContent(app())
	w.ShowAndRun()
}
