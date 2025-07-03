package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func new_start_page(start_func func()) *fyne.Container {
	title := widget.NewLabel("4Go10")
	title.SizeName = theme.SizeNameHeadingText
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	start_btn := widget.NewButton("Start", func() { go start_func() })

	high_score := widget.NewLabel("Highest Score - " + strconv.Itoa(readHighScore()))
	high_score.Alignment = fyne.TextAlignCenter

	page := container.New(
		layout.NewGridLayoutWithRows(5),
		&widget.Label{},
		&widget.Label{},
		container.New(layout.NewCenterLayout(), title),
		container.New(layout.NewCenterLayout(), start_btn),
		container.New(layout.NewCenterLayout(), high_score),
	)

	return page
}
