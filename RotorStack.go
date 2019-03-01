package main

import "fmt"
import "strconv"

type RotorStack struct {
	rotors [3]Rotor
}

func (rs *RotorStack) initStack(rotors [3]int, offsets [3]int) {
	 for index := 0; index < 3; index++ {
	 	rs.rotors[index] = rotorForNumber(rotors[index], offsets[index])
	}
}

func (rs *RotorStack) incrementOffset() {
	if rs.rotors[2].incrementOffset() {
		if rs.rotors[1].incrementOffset() {
			rs.rotors[0].incrementOffset()
		}
	}
}

func (rs RotorStack) encodeLeftIndex(index int) int {
	firstResult := rs.rotors[0].resultFromLeftIndex(index)
	secondResult := rs.rotors[1].resultFromLeftIndex(firstResult)
	thirdResult := rs.rotors[2].resultFromLeftIndex(secondResult)
	fmt.Println("Left Encode: " + strconv.Itoa(index))
	fmt.Println(strconv.Itoa(firstResult) + " " + strconv.Itoa(secondResult) + " " + strconv.Itoa(thirdResult))
	return rs.rotors[2].resultFromLeftIndex(rs.rotors[1].resultFromLeftIndex(rs.rotors[0].resultFromLeftIndex(index)))
}

func (rs RotorStack) encodeRightIndex(index int) int {
	firstResult := rs.rotors[2].resultFromRightIndex(index)
	secondResult := rs.rotors[1].resultFromRightIndex(firstResult)
	thirdResult := rs.rotors[0].resultFromRightIndex(secondResult)
	fmt.Println("Right Encode: " + strconv.Itoa(index))
	fmt.Println(strconv.Itoa(firstResult) + " " + strconv.Itoa(secondResult) + " " + strconv.Itoa(thirdResult))
	return rs.rotors[0].resultFromRightIndex(rs.rotors[1].resultFromRightIndex(rs.rotors[2].resultFromRightIndex(index)))
}