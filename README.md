# Brainfuck Go Interpreter

## Description
This library interpret brainfuck commands, and also you can add your custom commands and functions.

## How To Use
Applying commands
```
    brainFuck := brainfuck.New(instructions)
    brainFuck.ApplyCommands()
```
Add custom command
```
    brainFuck := brainfuck.New(instructions)
    brainFuck.AddCommand('$', func(brainfuck *brainfuck.BrainFuck) {
	    val := brainFuck.Cells[brainFuck.CellPointer]
	    brainFuck.Cells[brainFuck.CellPointer] = val * val
    })
```
