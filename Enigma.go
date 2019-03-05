package main

import (
"fmt"
)


func main() {
	fmt.Println("Hello, playground")

	var stack RotorReflectorStack
	stack.initStack()

	var plugBoard PlugBoard
	plugBoard.pairings = make(map[string]string)
	plugBoard.pair("A", "G")
	plugBoard.pair("B", "D")
	plugBoard.pair("E", "F")
	plugBoard.pair("H", "J")
	plugBoard.pair("K", "Y")
	plugBoard.pair("L", "M")
	plugBoard.pair("N", "P")
	plugBoard.pair("O", "C")
	plugBoard.pair("R", "T")
	plugBoard.pair("U", "I")

	fmt.Println(fullEncode("ILHGNTPCOO", plugBoard, &stack))

}

func fullEncode(message string, plugBoard PlugBoard, stack *RotorReflectorStack) string {
	var output string
	for i := 0; i < len(message); i++ {
		stack.rotorStack.incrementOffset()

		var inProgress = intForChar(plugBoard.valueForInput(string(message[i])))
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