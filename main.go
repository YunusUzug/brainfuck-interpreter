package main

import (
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
		brainFuck := New(instructions)
		brainFuck.ApplyCommands()
		err = brainFuck.AddCommand('$', func(brainfuck *BrainFuck) {
			val := brainFuck.Cells[brainFuck.CellPointer]
			brainFuck.Cells[brainFuck.CellPointer] = val * val
		})
		if err != nil {
			fmt.Printf("an error occured when add custom command: %v", err)
		}
	}
}
