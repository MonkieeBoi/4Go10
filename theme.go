package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type nordTheme struct{}

var colorNord0 = color.NRGBA{R: 46, G: 52, B: 64, A: 255}     // #2e3440
var colorNord1 = color.NRGBA{R: 59, G: 66, B: 82, A: 255}     // #3b4252
var colorNord2 = color.NRGBA{R: 67, G: 76, B: 94, A: 255}     // #434c5e
var colorNord3 = color.NRGBA{R: 76, G: 86, B: 106, A: 255}    // #4c566a
var colorNord4 = color.NRGBA{R: 216, G: 222, B: 233, A: 255}  // #d8dee9
var colorNord5 = color.NRGBA{R: 229, G: 233, B: 240, A: 255}  // #e5e9f0
var colorNord6 = color.NRGBA{R: 236, G: 239, B: 244, A: 255}  // #eceff4
var colorNord7 = color.NRGBA{R: 143, G: 188, B: 187, A: 255}  // #8fbcbb
var colorNord8 = color.NRGBA{R: 136, G: 192, B: 208, A: 255}  // #88c0d0
var colorNord9 = color.NRGBA{R: 129, G: 161, B: 193, A: 255}  // #81a1c1
var colorNord10 = color.NRGBA{R: 94, G: 129, B: 172, A: 255}  // #5e81ac
var colorNord11 = color.NRGBA{R: 191, G: 97, B: 106, A: 255}  // #bf616a
var colorNord12 = color.NRGBA{R: 208, G: 135, B: 112, A: 255} // #d08770
var colorNord13 = color.NRGBA{R: 235, G: 203, B: 139, A: 255} // #ebcb8b
var colorNord14 = color.NRGBA{R: 163, G: 190, B: 140, A: 255} // #a3be8c
var colorNord15 = color.NRGBA{R: 180, G: 142, B: 173, A: 255} // #b48ead

func (m nordTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark {
		switch name {
		case theme.ColorNamePrimary:
			return colorNord10
		case theme.ColorNameBackground:
			return colorNord0
		case theme.ColorNameButton:
			return colorNord1
		case theme.ColorNameInputBackground:
			return colorNord1
		case theme.ColorNameInputBorder:
			return colorNord2
		case theme.ColorNameFocus:
			return colorNord9
		}
	} else {
		switch name {
		case theme.ColorNamePrimary:
			return colorNord9
		case theme.ColorNameBackground:
			return colorNord6
		case theme.ColorNameButton:
			return colorNord5
		case theme.ColorNameInputBackground:
			return colorNord5
		case theme.ColorNameInputBorder:
			return colorNord4
		case theme.ColorNameFocus:
			return colorNord4
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m nordTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m nordTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m nordTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
