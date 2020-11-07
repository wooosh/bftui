package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var i *interpreter
var consoleText *tview.TextView
var consoleInput *tview.InputField

func createUI(i *interpreter) {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	programTape := newTape("Program Tape", instTape(&i.ip, &i.inst))
	dataTape := newTape("Data Tape", dataTape(&i.dp, &i.data))
	progOutput := newPrimitive("Program Output")
	progInput := newPrimitive("Program Input")
	consoleText = tview.NewTextView()
	consoleInput = tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorReset).
		SetDoneFunc(handleConsoleInput)

	grid := tview.NewGrid().
		SetRows(2, 2, 0, 1).
		SetColumns(0, 0, 0).
		SetBorders(true).
		AddItem(programTape, 0, 0, 1, 3, 0, 0, false).
		AddItem(dataTape, 1, 0, 1, 3, 0, 0, false).
		AddItem(progOutput, 2, 0, 2, 1, 0, 0, false).
		AddItem(progInput, 2, 1, 2, 1, 0, 0, false).
		AddItem(consoleText, 2, 2, 1, 1, 0, 0, false).
		AddItem(consoleInput, 3, 2, 1, 1, 0, 0, true)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one brainfuck file as an argument")
		os.Exit(1)
	}
	i = NewInterpreter()
	createUI(i)
}
