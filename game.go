package main

import (
	"math/rand/v2"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

func end_screen(page *fyne.Container, res_chan chan bool, num string, guess string) {
	page.Layout = layout.NewGridLayoutWithRows(9)
	page.Add(&widget.Label{})
	page.Add(&widget.Label{})

	l1 := widget.NewLabel("Number")
	l2 := widget.NewLabel(num)
	l3 := widget.NewLabel("Guessed")
	l4 := widget.NewLabel(guess)
	l5 := widget.NewButton("Ok", func() { res_chan <- false })

	l1.Alignment = fyne.TextAlignCenter
	l2.Alignment = fyne.TextAlignCenter
	l3.Alignment = fyne.TextAlignCenter
	l4.Alignment = fyne.TextAlignCenter

	l1.SizeName = theme.SizeNameSubHeadingText
	l2.SizeName = theme.SizeNameHeadingText
	l3.SizeName = theme.SizeNameSubHeadingText
	l4.SizeName = theme.SizeNameHeadingText

	l1.TextStyle = fyne.TextStyle{Bold: true}
	l3.TextStyle = fyne.TextStyle{Bold: true}

	page.Add(l1)
	page.Add(l2)
	page.Add(l3)
	page.Add(l4)
	page.Add(centered(l5))
	page.Refresh()
}

func addNum(boundNum binding.String, num string) func() {
	return func() {
		s, err := boundNum.Get()
		if err == nil {
			boundNum.Set(s + num)
		}
	}
}

func new_game_screen(digits int, res_chan chan bool) *fyne.Container {
	page := container.NewVBox()

	go func() {
		num := randNum(digits)
		numLabel := widget.NewLabel(num)
		numLabel.Alignment = fyne.TextAlignCenter
		numLabel.SizeName = theme.SizeNameHeadingText
		numLabel.Wrapping = fyne.TextWrapWord
		timer, done := timerWidget(calc_time(digits))

		fyne.DoAndWait(func() {
			page.Add(layout.NewSpacer())
			page.Add(stretched(numLabel))
			page.Add(layout.NewSpacer())
			page.Add(timer)
		})

		<-done

		boundNum := binding.NewString()
		boundLabel := widget.NewLabelWithData(boundNum)
		boundLabel.Alignment = fyne.TextAlignCenter
		boundLabel.SizeName = theme.SizeNameHeadingText
		boundLabel.Wrapping = fyne.TextWrapWord
		numPad := container.NewGridWithColumns(3)
		for i := 1; i <= 9; i++ {
			si := strconv.Itoa(i)
			button := stretched(widget.NewButton(si, addNum(boundNum, si)))
			numPad.Add(button)
		}

		backspace := func() {
			s, err := boundNum.Get()
			if err == nil && s != "" {
				runes := []rune(s)
				boundNum.Set(string(runes[:len(runes)-1]))
			}
		}
		numPad.Add(stretched(widget.NewButton("⌫", backspace)))
		numPad.Add(stretched(widget.NewButton("0", addNum(boundNum, "0"))))

		enter := func() {
			page.RemoveAll()
			w.Canvas().SetOnTypedKey(nil)
			page.Refresh()
			s, err := boundNum.Get()
			if s == num || err != nil {
				res_chan <- true
			} else {
				end_screen(page, res_chan, num, s)
			}
		}
		numPad.Add(stretched(widget.NewButton("✓", enter)))

		w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
			switch ke.Name {
			case "Return":
				enter()
			case "BackSpace":
				backspace()
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				s, err := boundNum.Get()
				if err == nil {
					boundNum.Set(s + string(ke.Name))
				}
			}
		})

		fyne.DoAndWait(func() {
			page.RemoveAll()
			page.Layout = layout.NewGridLayoutWithRows(2)
			page.Add(container.NewVBox(layout.NewSpacer(), boundLabel, layout.NewSpacer()))
			page.Add(numPad)
		})
	}()

	return page
}
