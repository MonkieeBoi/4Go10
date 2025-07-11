package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func centered(widget fyne.Widget) *fyne.Container {
	return container.New(layout.NewCenterLayout(), widget)
}
