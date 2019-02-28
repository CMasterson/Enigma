package main

type RotorStack struct {
	rotors [3]Rotor
}

func (rs *RotorStack) incrementOffset() {
	if rs.rotors[2].incrementOffset() {
		if rs.rotors[1].incrementOffset() {
			rs.rotors[0].incrementOffset()
		}
	}
}