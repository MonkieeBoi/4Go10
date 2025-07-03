package main

import (
	"math/rand/v2"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func randNum(digits int) string {
	d := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if digits == 1 {
		return d[rand.IntN(len(d))]
	}
	res := d[rand.IntN(len(d)-1)+1]

	for range digits - 1 {
		res += d[rand.IntN(len(d))]
	}

	return res
}

func calc_time(digits int) time.Duration {
	return time.Second * time.Duration(digits)
}

func new_game_screen(digits int, result_chan chan bool) *fyne.Container {
	page := container.NewVBox()

	go func() {
		num := randNum(digits)
		page.Add(widget.NewLabel(num))
		timer, done := timerWidget(calc_time(digits))
		page.Add(timer)
		<-done

		page.RemoveAll()
		entry := widget.NewEntry()
		page.Add(entry)
		entry.OnSubmitted = func(s string) {
			page.RemoveAll()
			if s != num {
				result_chan <- false
			} else {
				result_chan <- true
			}
		}
	}()

	return page
}
