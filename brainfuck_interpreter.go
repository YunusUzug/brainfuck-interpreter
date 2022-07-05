/*
>	increment pointer
<	decrement pointer
+	increment value at pointer
-	decrement value at pointer
[	begin loop (continues while value at pointer is non-zero)
]	end loop
,	read one character from input into value at pointer
.	print value at pointer to output as a character
and also user can define
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	fn func(brainfuck *BrainFuck)

	BrainFuck struct {
		cellSize    int
		CellPointer int
		Cells       []byte
		instruction string
		commandList []interface{}
	}

	loopCommand struct {
		instruction string
		commandList []interface{}
	}
)

var commandMap = map[byte]fn{
	'>': incrementPointer,
	'<': decrementPointer,
	'+': incrementPointerValue,
	'-': decrementPointerValue,
	',': readOneCharacterFromInputIntoValueAtPointer,
	'.': printValueAtPointer,
}

func (b *BrainFuck) AddCommand(command byte, function fn) (error, bool) {
	if _, ok := commandMap[command]; ok {
		return fmt.Errorf("command is already exist"), false
	}
	commandMap[command] = function
	return nil, true
}

func (b *BrainFuck) ApplyCommands() {
	b.Cells = make([]byte, 30_000)

	prepareInstructionList(b, nil, b.instruction)
	for _, val := range b.commandList {
		applyCommand(val, b)
	}
}

func New(instructions string) *BrainFuck {
	brainFuck := &BrainFuck{
		cellSize:    30_000,
		CellPointer: 0,
		Cells:       make([]byte, 30_000),
		instruction: instructions,
		commandList: make([]interface{}, 0),
	}
	return brainFuck
}

/*
	Decide the type of val. If the val type is LoopCommand then applies loop operations.
	In case of val is default apply the related function.
*/
func applyCommand(val interface{}, b *BrainFuck) {
	switch val.(type) {
	case *loopCommand:
		val.(*loopCommand).applyLoopCommands(b)
	default:
		if fn, ok := commandMap[val.(byte)]; ok {
			fn(b)
		}
	}
}

/*
	Apply loop commands until the cell value which is pointed by cellPointer is zero.
*/
func (l loopCommand) applyLoopCommands(brainFuck *BrainFuck) {
	for brainFuck.Cells[brainFuck.CellPointer] > 0 {
		for _, val := range l.commandList {
			applyCommand(val, brainFuck)
		}
	}
}

/*
>	increment pointer
*/
func incrementPointer(brainFuck *BrainFuck) {
	if brainFuck.CellPointer == brainFuck.cellSize-1 {
		brainFuck.CellPointer = 0
	} else {
		brainFuck.CellPointer += 1
	}
}

/*
<	decrement pointer
*/
func decrementPointer(brainFuck *BrainFuck) {
	if brainFuck.CellPointer == 0 {
		brainFuck.CellPointer = brainFuck.cellSize - 1
	} else {
		brainFuck.CellPointer -= 1
	}
}

/*
	+	increment value at pointer
*/
func incrementPointerValue(brainFuck *BrainFuck) {
	brainFuck.Cells[brainFuck.CellPointer] += 1
}

/*
	-	decrement value at pointer
*/
func decrementPointerValue(brainFuck *BrainFuck) {
	brainFuck.Cells[brainFuck.CellPointer] -= 1
}

/*
	.	print value at pointer to output as a character
*/
func printValueAtPointer(brainFuck *BrainFuck) {
	fmt.Printf("%c", rune(brainFuck.Cells[brainFuck.CellPointer]))
}

/*
	,	read one character from input into value at pointer
*/
func readOneCharacterFromInputIntoValueAtPointer(brainFuck *BrainFuck) {
	fmt.Print("input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	if err != nil {
		fmt.Printf("an error occured when read input: %v", err)
		return
	}
	brainFuck.Cells[brainFuck.CellPointer] = byte(input)
}

/*
	Prepares the instruction list by commands. If detect any loops then calls itself by subCommands.
*/
func prepareInstructionList(brainFuck *BrainFuck, lc *loopCommand, instruction string) {
	for i := 0; i < len(instruction); i++ {
		char := instruction[i]
		if char == '[' {
			loopStartIndex := i + 1
			loopEndIndex := loopStartIndex
			leftLoopCharCount := 0

			for j := loopStartIndex; j < len(instruction); j++ {
				subChar := instruction[j]
				if subChar == '[' {
					leftLoopCharCount++
				} else if subChar == ']' {
					if leftLoopCharCount == 0 {
						loopEndIndex = j
						break
					} else {
						leftLoopCharCount--
					}
				}
			}

			loopInstruction := instruction[loopStartIndex:loopEndIndex]
			i = loopEndIndex
			subLoopCommand := &loopCommand{
				instruction: loopInstruction,
				commandList: make([]interface{}, 0),
			}
			if lc == nil {
				brainFuck.commandList = append(brainFuck.commandList, subLoopCommand)
			} else {
				lc.commandList = append(lc.commandList, subLoopCommand)
			}
			prepareInstructionList(brainFuck, subLoopCommand, loopInstruction)
		} else if char == ']' {
			return
		} else {
			if lc != nil {
				lc.commandList = append(lc.commandList, char)
			} else {
				brainFuck.commandList = append(brainFuck.commandList, char)
			}
		}
	}
}
