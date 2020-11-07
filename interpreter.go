package main

import (
	"io/ioutil"
)

type interpreter struct {
	ip   int
	inst []byte

	dp   int
	data []byte

	/*
	  output []byte
	  input []byte
	*/
}

func NewInterpreter() *interpreter {
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
	i.data = make([]byte, 1)
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
		//fmt.Print(string(i.data[i.dp]))
	case ',':
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
