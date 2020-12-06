package five

import (
	"sort"
	"strings"
)

/**
Search example for input -> F B F B B F F
NOTE: we care about the upper bound only

127			0000000 - 1111111 0 - 127
64		F	0000000 - 0111111 0 - 63		clear bit on 0
32		B 0100000 - 0111111 32 - 63		no clear
16		F 0100000 - 0101111 32 - 47		clear bit on 2
8			B 0101000 - 0101111 40 - 47		no clear
4			B 0101100 - 0101111 44 - 47		no clear
2			F 0101100 - 0101101 44 - 45		clear bit on 5
1			F 0101100 44
*/

func bSearch(input string, flipper rune) int {
	// Start with length of input number of 1 bits (1111111 - 127)
	result := (1 << len(input)) - 1

	for idx, splitDirection := range input {
		// We only clear on "halving" directions F / L
		if splitDirection == flipper {
			result ^= 1 << (len(input) - (idx + 1))
		}
	}

	return result
}

func getSeatID(input string) int {
	row := bSearch(input[:len(input)-3], 'F')
	column := bSearch(input[len(input)-3:], 'L')

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
