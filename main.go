package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("4Go10")

	timer := widget.NewProgressBar()
	timer.TextFormatter = func() string {return ""}

	go func() {
		secs := 10.0
		for i := 1.0; i >= 0; i -= 0.001 / secs {
			time.Sleep(time.Millisecond * 1)
			fyne.Do(func() {
				timer.SetValue(i)
			})
		}
	}()

	w.SetContent(container.NewVBox(timer))
	w.ShowAndRun()
}
