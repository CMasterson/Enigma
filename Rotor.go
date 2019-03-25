package main

/*
Num | Char Sequence              | Turnover | Numerical Sequence (Adjusted for Turnover)
  I	| EKMFLGDQVZNTOWYHXUSPAIBRCJ |	  Q	    | 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9, 4, 10, 12, 5, 11, 6, 3
 II	| AJDKSIRUXBLHWTMCQGZNPYFVOE |	  E	    | 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4, 0, 9, 3, 10
III	| BDFHJLCPRTXVZNYEIWGAKMUSQO |	  V	    | 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14, 1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23
 IV	| ESOVPZJAYQUIRHXLNFTGKDCMWB |	  J	    |
  V	| VZBRGITYUPSDNHLXAWMJQOFECK |	  Z	    |
*/

 var rotorSequences = [3][26]int{ 
 	[26]int{16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9, 4, 10, 12, 5, 11, 6, 3},
    [26]int{18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4, 0, 9, 3, 10},
    [26]int{21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14, 1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23},
}

type Rotor struct {
	sequence [26]int
	currentOffset int
}

func rotorForNumber(rotorNumber int, offset int) Rotor {
	return Rotor {rotorSequences[rotorNumber], offset}
}

func (r Rotor) resultFromLeftIndex(leftIndex int) int {
	adjustedIndex := (leftIndex + r.currentOffset) % len(r.sequence)
	rightSideIndex := r.sequence[adjustedIndex]
	adjustedRightSideIndex := rightSideIndex - r.currentOffset
	if adjustedRightSideIndex < 0 {
		adjustedRightSideIndex += len(r.sequence)
	}
	return adjustedRightSideIndex
}

func (r Rotor) resultFromRightIndex(rightIndex int) int {
	adjustedIndex := (rightIndex + r.currentOffset) % len(r.sequence)

	for leftIndex, rightValue := range r.sequence {
		if rightValue == adjustedIndex {
			adjustedLeftIndex := leftIndex - r.currentOffset
			if adjustedLeftIndex < 0 {
				adjustedLeftIndex += len(r.sequence)
			}

			return adjustedLeftIndex
		}
	}

	return 0
}


func (r *Rotor) incrementOffset() bool {
	r.currentOffset = (r.currentOffset + 1) % len(r.sequence)
	return r.currentOffset == 0
}

