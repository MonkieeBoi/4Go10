package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func mainApp() *fyne.Container {
	started := make(chan bool, 1)
	res := make(chan bool, 1)
	win := container.New(layout.NewPaddedLayout())
	start_func := func() { started <- true }

	main_loop := func() {
		for {
			start_page := new_start_page(start_func)
			win.RemoveAll()
			win.Add(start_page)
			<-started

			digits := 1
			for {
				win.RemoveAll()
				win.Add(new_game_screen(digits, res))
				won := <-res
				if !won {
					updateHighScore(digits)
					break
				}
				digits++
			}
		}
	}

	go main_loop()

	return win
}
