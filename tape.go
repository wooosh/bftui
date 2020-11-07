package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func instTape(ip *int, inst *[]byte) func() (int, []rune) {
	return func() (int, []rune) {
		return *ip, []rune(string(*inst))
	}
}

func dataTape(dp *int, data *[]byte) func() (int, []rune) {
	return func() (int, []rune) {
		var buf string
		var index int
		for i, c := range *data {
			if i == *dp {
				index = len(buf)
			}
			buf += strconv.Itoa(int(c)) + " "
		}
		return index, []rune(buf)
	}
}

// @Todo: should take a render function, for hex, int, and char mode
func newTape(name string, render func() (int, []rune)) *tview.Box {
	return tview.NewBox().
		SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {

			// Draw a horizontal line across the middle of the box.
			centerY := y + height/2
			index, tape := render()
			tapeStartX := (width / 2) - index + x
			for i, inst := range tape {
				if i+tapeStartX >= x+width {
					break
				}
				style := tcell.StyleDefault
				if i == index {
					style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
				}
				screen.SetContent(tapeStartX+i, centerY, rune(inst), nil, style)
			}
			// Write some text along the horizontal line.
			tview.Print(screen, name, x, y, width, tview.AlignCenter, tcell.ColorWhite)
			//tview.Print(screen, string(*tape), x, y+1, width-2, tview.AlignCenter, tcell.ColorReset)

			// Space for other content.
			return x + 1, 2, width - 2, height - (centerY + 1 - y)
		})
}
