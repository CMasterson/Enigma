package main

import (
"fmt"
)

func incrementRotors(rotors *[3]Rotor) {
	if rotors[0].incrementOffset() {
		if rotors[1].incrementOffset() {
			rotors[2].incrementOffset()
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	
	rotors := createRotors()
	fmt.Println(rotors[0].resultFromLeftIndex(0))
	fmt.Println(rotors[0].resultFromRightIndex(0))
	fmt.Println(rotors[0].incrementOffset())
	fmt.Println(rotors[0].resultFromLeftIndex(0))
}

func createRotors() [3]Rotor {
	return [3]Rotor {Rotor{[...]int{24, 0, 14, 6, 16, 1, 15, 4, 5, 2, 3, 9, 18, 17, 8, 22, 20, 13, 25, 19, 23, 10, 21, 12, 7, 11}, 5}, Rotor{[...]int{24, 0, 14, 6, 16, 1, 15, 4, 5, 2, 3, 9, 18, 17, 8, 22, 20, 13, 25, 19, 23, 10, 21, 12, 7, 11}, 0},  Rotor{[...]int{24, 0, 14, 6, 16, 1, 15, 4, 5, 2, 3, 9, 18, 17, 8, 22, 20, 13, 25, 19, 23, 10, 21, 12, 7, 11}, 0} }
}