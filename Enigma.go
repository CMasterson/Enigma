package main

import (
"fmt"
"strings"
)


func main() {
	fmt.Println("Hello, playground")

	var stack RotorReflectorStack
	stack.initStack()

	var plugBoard PlugBoard
	plugBoard.pairings = make(map[string]string)
	plugBoard.pair("B", "Q")
	plugBoard.pair("C", "R")
	plugBoard.pair("D", "I")
	plugBoard.pair("E", "J")
	plugBoard.pair("K", "W")
	plugBoard.pair("M", "T")
	plugBoard.pair("O", "S")
	plugBoard.pair("P", "X")
	plugBoard.pair("U", "Z")
	plugBoard.pair("G", "H")

	fmt.Println(fullEncode("Lorem ipsum", plugBoard, &stack))

}

func fullEncode(message string, plugBoard PlugBoard, stack *RotorReflectorStack) string {
	var output string

	var formattedMessage = strings.ToUpper(message)

	for i := 0; i < len(formattedMessage); i++ {
		var char = string(formattedMessage[i])

		if char == " " {
			output = output + " "
			continue
		}

		stack.rotorStack.incrementOffset()

		var inProgress = intForChar(plugBoard.valueForInput(char))
		inProgress = stack.code(inProgress)
		var codedChar = plugBoard.valueForInput(charForInt(inProgress))

		output = output + codedChar
	}

	return output
}

func charForInt(code int) string {
	return string(byte(code + 65))
}

func intForChar(char string) int {
	return int(byte(char[0])) - 65
}