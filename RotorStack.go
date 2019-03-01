package main

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
	return rs.rotors[2].resultFromLeftIndex(rs.rotors[1].resultFromLeftIndex(rs.rotors[0].resultFromLeftIndex(index)))
}

func (rs RotorStack) encodeRightIndex(index int) int {
	return rs.rotors[0].resultFromRightIndex(rs.rotors[1].resultFromRightIndex(rs.rotors[2].resultFromRightIndex(index)))
}