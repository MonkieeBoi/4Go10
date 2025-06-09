package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func app() *fyne.Container {
	title := widget.NewLabel("4Go10")
	title.SizeName = theme.SizeNameHeadingText
	title.Alignment = fyne.TextAlignCenter

	win := container.New(layout.NewVBoxLayout(), title)

	return win
}
