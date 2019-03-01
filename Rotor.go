package main
// import "fmt"
/*
Num | Char Sequence              | Turnover | Numerical Sequence (Adjusted for Turnover)
  I	| EKMFLGDQVZNTOWYHXUSPAIBRCJ |	  Q	    | 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9, 4, 10, 12, 5, 11, 6, 3
 II	| AJDKSIRUXBLHWTMCQGZNPYFVOE |	  E	    | 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4, 0, 9, 3, 10
III	| BDFHJLCPRTXVZNYEIWGAKMUSQO |	  V	    |
 IV	| ESOVPZJAYQUIRHXLNFTGKDCMWB |	  J	    |
  V	| VZBRGITYUPSDNHLXAWMJQOFECK |	  Z	    |
*/

 var rotorSequences = [3][26]int{ 
 	[26]int{16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9, 4, 10, 12, 5, 11, 6, 3},
    [26]int{18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4, 0, 9, 3, 10},
    [26]int{16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9, 4, 10, 12, 5, 11, 6, 3},
}

type Rotor struct {
	sequence [26]int
	currentOffset int
}

func rotorForNumber(rotorNumber int, offset int) Rotor {
	return Rotor {rotorSequences[rotorNumber], offset}
}

func (r Rotor) resultFromLeftIndex(leftIndex int) int {
	// fmt.Println("========Encoding=======")
	// fmt.Println(leftIndex)
	adjustedIndex := (leftIndex + r.currentOffset) % len(r.sequence)
	// fmt.Println("Adjusted Index")
	// fmt.Println(adjustedIndex)
	rightSideIndex := r.sequence[adjustedIndex] 
	// fmt.Println("Pre-adjusted Output")
	// fmt.Println(rightSideIndex)
	adjustedRightSideIndex := rightSideIndex - r.currentOffset
	if adjustedRightSideIndex < 0 {
		adjustedRightSideIndex += len(r.sequence)
	}
	// fmt.Println("Adjusted Output")
	// fmt.Println(adjustedRightSideIndex)
	// fmt.Println("=====END Left Encode=====")
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

