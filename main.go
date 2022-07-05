package main

import (
	brainfuck "brainfuck-interpreter/application"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		file, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("an error occured when read file: %v", err)
			return
		}

		instructions := string(file)
		brainFuck := brainfuck.New(instructions)
		brainFuck.ApplyCommands()
		brainFuck.AddCommand('$', func(brainfuck *brainfuck.BrainFuck) {
			val := brainFuck.Cells[brainFuck.CellPointer]
			brainFuck.Cells[brainFuck.CellPointer] = val * val
		})
	}
}
