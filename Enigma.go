package main

import (
"fmt"
)


func main() {
	fmt.Println("Hello, playground")

	var stack RotorReflectorStack
	stack.initStack()

	fmt.Println(stack.code(0))
	fmt.Println(stack.code(21))
}