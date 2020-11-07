package main

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// @Todo: make sure interpreter is ready
// @Todo: input command
// @Todo: breakpoints
// @Todo: reload file
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
				go i.step()
			} else if len(cmd) == 2 {
				num, err := strconv.Atoi(cmd[1])
				if err != nil {
					console += "unable to parse number of instructions"
				} else {
					for n := 0; n < num; n++ {
						go i.step()
					}
				}
			} else {
				console += "expected zero arguments to step\n"
			}
		case "r", "run":
			if len(cmd) != 1 {
				console += "expected zero arguments to run\n"
			} else {
				go i.run()
			}
		case "numfmt":
			if len(cmd) != 3 {
				console += "expected three arguments to numfmt\n"
			} else {
				if cmd[2] != "decimal" && cmd[2] != "hex" && cmd[2] != "mixed" && cmd[2] != "ascii" {
					console += "invalid mode type '" + cmd[2] + "'\n"
				} else {
					if cmd[1] == "tape" {
						dataTapeDisplayMode = cmd[2]
					} else if cmd[1] == "output" {
						outputDisplayMode = cmd[2]
					} else {
						console += "expected first argument to be 'tape' or 'output'\n"
					}
				}
			}
		case "input":
			if len(cmd) != 2 {
				console += "expected one argument to input\n"
			}
			for _, b := range cmd[1] {
				i.input <- byte(b)
			}
		case "q", "quit":
			app.Stop()
		default:
			console += "available commands:\n" +
				"  o|open <file> - opens a file\n" +
				"  r|run - runs the opened file\n" +
				"  s|step [n] - executes n instructions (default 1)\n" +
				"  numfmt <tape|output> <decimal|hex|mixed|ascii> - changes how numbers are displayed in the data tape or program output. Mixed falls back to decimal for non-printable characters\n" +
				"  input <text> - sends input to the brainfuck interpreter\n" +
				"  q|quit - quits the program\n"
		}

		consoleText.SetText(console)
	}
}
