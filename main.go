package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	decimal = "decimal"
	hex     = "hex"
	ascii   = "ascii"
	mixed   = "mixed"
)

func convertByteToString(b byte, mode string) string {
	switch mode {
	case decimal:
		return strconv.Itoa(int(b))
	case hex:
		return fmt.Sprintf("%X", int(b))
	case ascii:
		if unicode.IsPrint(rune(b)) {
			return string(b)
		} else {
			return "."
		}
	case mixed:
		if unicode.IsPrint(rune(b)) {
			return "'" + string(b) + "'"
		} else {
			return strconv.Itoa(int(b))
		}
	}
	panic("how")
}

var i *interpreter
var consoleText *tview.TextView
var consoleInput *tview.InputField
var progOutput *tview.TextView
var app *tview.Application

var dataTapeDisplayMode string = mixed
var outputDisplayMode string = ascii

func createUI(i *interpreter) {
	programTape := newTape(instTape(&i.ip, &i.inst))
	dataTape := newTape(dataTape(&i.dp, &i.data))
	progOutput = tview.NewTextView().
		SetText("Program Output (ascii)\n")
	consoleText = tview.NewTextView()
	consoleInput = tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorReset).
		SetDoneFunc(handleConsoleInput)

	grid := tview.NewGrid().
		SetRows(2, 2, 0, 1).
		SetColumns(0, 0).
		SetBorders(true).
		AddItem(programTape, 0, 0, 1, 2, 0, 0, false).
		AddItem(dataTape, 1, 0, 1, 2, 0, 0, false).
		AddItem(progOutput, 2, 0, 2, 1, 0, 0, false).
		AddItem(consoleText, 2, 1, 1, 1, 0, 0, false).
		AddItem(consoleInput, 3, 1, 1, 1, 0, 0, true)

	app = tview.NewApplication()
	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}

func main() {
	/*
		if len(os.Args) != 2 {
			fmt.Println("Please provide exactly one brainfuck file as an argument")
			os.Exit(1)
		}*/
	i = newInterpreter()
	createUI(i)
}
