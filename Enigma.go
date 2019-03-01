package main

import (
"fmt"
)


func main() {
	fmt.Println("Hello, playground")

	var rotorStack RotorStack
	rotorStack.initStack([3]int{0,1,2}, [3]int{0,0,0})

	fmt.Println(rotorStack.encodeLeftIndex(1))
	fmt.Println(rotorStack.encodeRightIndex(19))
}