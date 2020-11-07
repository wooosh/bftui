package main

import (
	"io/ioutil"
)

type interpreter struct {
	ip   int
	inst []byte

	dp   int
	data []byte

	output []byte
	input  chan byte
	// @Todo: make output a channel
	//input []byte
}

func newInterpreter() *interpreter {
	var i interpreter
	return &i
}

// returns true if it was able to open the file
func (i *interpreter) openFile(filename string) bool {
	var err error
	i.inst, err = ioutil.ReadFile(filename)
	if err != nil {
		return false
	}
	i.ip = -1
	i.dp = 0
	i.data = make([]byte, 1)
	i.output = make([]byte, 0)
	i.input = make(chan byte)
	return true
}

// returns whether or not it could step
func (i *interpreter) step() bool {
	// @Todo: how to handle end of program?
	if i.ip+1 == len(i.inst) {
		return false
	}
	i.ip++
	switch i.inst[i.ip] {
	case '>':
		// make sure we have space
		if i.dp+1 == len(i.data) {
			i.data = append(i.data, 0)
		}
		i.dp++
	case '<':
		if i.dp > 0 {
			i.dp--
		}
	case '+':
		i.data[i.dp]++
	case '-':
		i.data[i.dp]--
	case '.':
		i.output = append(i.output, i.data[i.dp])
		var buf string
		for _, b := range i.output {
			buf += convertByteToString(b, outputDisplayMode)
		}
		progOutput.SetText(buf)
	// Signify we are waiting for input
	case ',':
		i.data[i.dp] = <-i.input
	case '[':
		if i.data[i.dp] == 0 {
			nests := 1
			for {
				i.ip++ // @Todo: bounds checking
				if i.inst[i.ip] == '[' {
					nests++
				} else if i.inst[i.ip] == ']' {
					nests--
				}
				if nests == 0 {
					break
				}
			}
			//fmt.Println("No closing bracket was found for opening bracket at pos", i.dp)
		}
	case ']':
		if i.data[i.dp] != 0 {
			nests := 1
			for {
				i.ip-- // @Todo: bounds checking
				if i.inst[i.ip] == ']' {
					nests++
				} else if i.inst[i.ip] == '[' {
					nests--
				}
				if nests == 0 {
					break
				}
			}
			// fmt.Println("No opening bracket was found for closing bracket at pos", i.dp)
		}
	default:
		return i.step()
	}
	return true
}

func (i *interpreter) run() {
	for i.step() {
	}
}
