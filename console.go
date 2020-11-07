package main

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func handleConsoleInput(key tcell.Key) {
	if key == tcell.KeyEnter {
		cmdText := consoleInput.GetText()
		cmd := strings.Split(cmdText, " ")
		consoleInput.SetText("")
		if len(cmd) == 0 {
			return
		}

		console := consoleText.GetText(true)
		console += "> " + cmdText + "\n"

		switch cmd[0] {
		case "o", "open":
			if len(cmd) != 2 {
				console += "expected one argument to open\n"
			} else if !i.openFile(cmd[1]) {
				console += "unable to open file '" + cmd[1] + "'\n"
			}
		case "s", "step":
			if len(cmd) == 1 {
				i.step()
			} else if len(cmd) == 2 {
				num, err := strconv.Atoi(cmd[1])
				if err != nil {
					console += "unable to parse number of instructions"
				} else {
					for n := 0; n < num; n++ {
						i.step()
					}
				}
			} else {
				console += "expected zero arguments to step\n"
			}
		case "r", "run":
			if len(cmd) != 1 {
				console += "expected zero arguments to run\n"
			} else {
				i.run()
			}
		case "q", "quit":

		default:
			console += "available commands:\n" +
				"  o|open <file> - opens a file\n" +
				"  r|run - runs the opened file\n" +
				"  s|step [n]- executes n instructions (default 1)\n" +
				"  q|quit - quits the program\n"
		}

		consoleText.SetText(console)
	}
}
