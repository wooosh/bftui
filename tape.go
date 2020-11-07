package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func instTape(ip *int, inst *[]byte) func() (string, int, int, []rune) {
	return func() (string, int, int, []rune) {
		return "Program Tape", *ip, 1, []rune(string(*inst))
	}
}

func dataTape(dp *int, data *[]byte) func() (string, int, int, []rune) {
	return func() (string, int, int, []rune) {
		var buf string
		var index int
		var width int
		for i, b := range *data {
			if i == *dp {
				index = len(buf)
			}
			buf += convertByteToString(b, dataTapeDisplayMode) + " "
			if i == *dp {
				width = len(buf) - index
			}
		}
		return "Data Tape (" + dataTapeDisplayMode + ")", index, width, []rune(buf)
	}
}

// @Todo: highlight whole number
// render is title, index, width of index, full rendered tape
func newTape(render func() (string, int, int, []rune)) *tview.Box {
	return tview.NewBox().
		SetDrawFunc(func(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {

			centerY := y + height/2
			title, index, iwidth, tape := render()
			tapeStartX := (width / 2) - index + x
			for i, inst := range tape {
				if i+tapeStartX >= x+width {
					break
				}
				style := tcell.StyleDefault
				if i >= index && i < index+iwidth {
					style = tcell.StyleDefault.Foreground(tcell.ColorYellow)
				}
				screen.SetContent(tapeStartX+i, centerY, rune(inst), nil, style)
			}
			// Write some text along the horizontal line.
			tview.Print(screen, title, x, y, width, tview.AlignCenter, tcell.ColorWhite)
			//tview.Print(screen, string(*tape), x, y+1, width-2, tview.AlignCenter, tcell.ColorReset)

			// Space for other content.
			return x + 1, 2, width - 2, height - (centerY + 1 - y)
		})
}
