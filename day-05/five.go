package five

import (
	"sort"
	"strings"
)

/**
Search example for input -> F B F B B F F

127		0000000 - 1111111 0 - 127
64	F 0000000 - 0111111 0 - 63
32	B 0100000 - 0111111 32 - 63
16	F 0100000 - 0101111 32 - 47
8 	B 0101000 - 0101111 40 - 47
4 	B 0101100 - 0101111 44 - 47
2 	F 0101100 - 0101101 44 - 45
1 	F 0101100 44
*/

func bSearch(input string, left rune, right rune) int {
	upperBound := (1 << len(input)) - 1
	spread := upperBound
	lowerBound := 0

	for _, splitDirection := range input {
		spread = spread >> 1
		if splitDirection == left {
			upperBound -= spread + 1
		} else if splitDirection == right {
			lowerBound += spread + 1
		}
	}

	return upperBound
}

func getSeatID(input string) int {
	row := bSearch(input[:len(input)-3], 'F', 'B')
	column := bSearch(input[len(input)-3:], 'L', 'R')

	return row*8 + column
}

func getLargestSeatID(input string) int {
	foundSolution := 0

	for _, seat := range strings.Split(input, "\n") {
		seatID := getSeatID(seat)
		if seatID > foundSolution {
			foundSolution = seatID
		}
	}

	return foundSolution
}

func getMySeat(input string) int {
	seats := strings.Split(input, "\n")
	seatIDs := make([]int, len(seats))

	for idx, seat := range seats {
		seatIDs[idx] = getSeatID(seat)
	}

	sort.Ints(seatIDs)
	middle := len(seatIDs) / 2
	startOffset := seatIDs[0]

	for idx := 0; idx < middle; idx++ {
		if seatIDs[middle+idx] != startOffset+middle+idx {
			return startOffset + middle + idx
		}
		if seatIDs[middle-idx-1] != startOffset+middle-idx-1 {
			return startOffset + middle - idx
		}
	}

	return 0
}
