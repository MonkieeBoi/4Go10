package main

import (
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := fyneApp.New()
	a.Settings().SetTheme(theme.DefaultTheme())
	w := a.NewWindow("4Go10")
	w.SetContent(app())
	w.ShowAndRun()
}
