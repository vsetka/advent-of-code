package five

import (
	"regexp"
	"sort"
	"strconv"
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

func getSeatID(input string) int {
	// In order to get the seat ID, we need to look at our input in binary
	// representation where F/L are 1s and B/R are 0s
	input = regexp.MustCompile(`F|L`).ReplaceAllString(input, "1")
	input = regexp.MustCompile(`B|R`).ReplaceAllString(input, "0")
	binaryInput, _ := strconv.ParseInt(input, 2, 16)

	// Once we have converted the search input into its binary representation
	// we can just flip the bits against the same size binary number with all bits set to 1
	// We can do this all in one go as this is equivalent to calculating rows and columns separately
	// and merging them effectively into one number by multiplying rows by 8 and adding to columns
	// which is the same as shifting the rows left by 3 and merging with columns, ie (row << 3) | column
	return int(binaryInput ^ ((1 << len(input)) - 1))
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
