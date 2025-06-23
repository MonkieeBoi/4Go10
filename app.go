package main

import (
	"math/rand/v2"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func randNum(digits int) string {
	d := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	res := ""

	for range digits {
		res += d[rand.IntN(len(d))]
	}

	return res
}

func app() *fyne.Container {
	title := widget.NewLabel("4Go10")
	title.SizeName = theme.SizeNameHeadingText
	title.Alignment = fyne.TextAlignCenter

	win := container.New(layout.NewVBoxLayout(), title)
	digits := 1

	main_loop := func() {
		win.RemoveAll()
		num := randNum(digits)
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
				digits++
			}
		}
	}

	start_btn := widget.NewButton("Start", func() { go main_loop() })
	win.Add(start_btn)

	return win
}
