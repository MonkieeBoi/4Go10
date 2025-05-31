package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func timerWidget(d time.Duration) (*widget.ProgressBar, chan bool) {
	timer := widget.NewProgressBar()
	timer.TextFormatter = func() string { return "" }
	done := make(chan bool, 1)

	go func() {
		secs := d.Seconds()
		for i := 1.0; i >= 0; i -= 0.001 / secs {
			time.Sleep(time.Millisecond * 1)
			fyne.Do(func() {
				timer.SetValue(i)
			})
		}
		done <- true
		close(done)
	}()
	return timer, done
}
