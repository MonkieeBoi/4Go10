package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func centered(canvas fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewCenterLayout(), canvas)
}

func stretched(canvas fyne.CanvasObject) *fyne.Container {
	return container.New(layout.NewStackLayout(), canvas)
}
