package five

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/**
Search example for input -> F B F B B F F

127			0000000 - 1111111 0 - 127
64		F	0000000 - 0111111 0 - 63		clear bit on 0
32		B 0100000 - 0111111 32 - 63		no clear
16		F 0100000 - 0101111 32 - 47		clear bit on 2
8			B 0101000 - 0101111 40 - 47		no clear
4			B 0101100 - 0101111 44 - 47		no clear
2			F 0101100 - 0101101 44 - 45		clear bit on 5
1			F 0101100 44 (basically our input, FBFBBFF if F is 0 and B is 1)

Turns out the input pattern describing the seat position in 2D space by doing binary
search with back/front + left/right instructions is just an obfuscated binary representation
of the seat ID. Each letter encodes either 1 or 0. The clue was in the the number of rows and seats
as well as the final encoding of the seat ID by multiplying rows by 8 and adding columns (ie -> (row << 3) | column).
*/
func getSeatID(input string) int {
	input = regexp.MustCompile(`F|L`).ReplaceAllString(input, "0")
	input = regexp.MustCompile(`B|R`).ReplaceAllString(input, "1")
	binaryInput, _ := strconv.ParseInt(input, 2, 16)

	return int(binaryInput)
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
			return startOffset + middle - idx - 1
		}
	}

	return 0
}
