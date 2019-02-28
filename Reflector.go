package main
/*
Name   | Char sequence
UKW-A  | EJMZALYXVBWFCRQUONTSPIKHGD  |  
UKW-B* | YRUHQSLDPXNGOKMIEBFZCWVJAT  |  24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19
UKW-C  | FVPJIAOYEDRZXWGCTKUQSBNMHL  |
*/

type Reflector struct {
	sequence [26]int
}

func (r Reflector) outputForInput(input int) int {
	return r.sequence[input]
}