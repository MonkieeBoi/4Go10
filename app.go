package main

import (
	"math/rand/v2"
	"strconv"
	"time"

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

	main_loop := func() {
		win.RemoveAll()
		num := strconv.Itoa(rand.IntN(100))
		win.Add(widget.NewLabel(num))
		timer, done := timerWidget(time.Duration(1 * time.Second))
		win.Add(timer)
		_ = <-done
		win.RemoveAll()
		entry := widget.NewEntry()
		win.Add(entry)
		entry.OnSubmitted = func(s string) {
			win.RemoveAll()
			if s != num {
				win.Add(widget.NewLabel("Wrong"))
			} else {
				win.Add(widget.NewLabel("Correct"))
			}
		}
	}

	start_btn := widget.NewButton("Start", func() { go main_loop() })
	win.Add(start_btn)

	return win
}
