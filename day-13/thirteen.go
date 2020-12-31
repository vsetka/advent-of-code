package thirteen

import (
	"strconv"
	"strings"
)

func parse(input string) (int, []int) {
	rows := strings.Split(input, "\n")
	timestamp, _ := strconv.Atoi(rows[0])
	buses := []int{}

	for _, slot := range strings.Split(rows[1], ",") {
		if slot != "x" {
			bus, _ := strconv.Atoi(slot)
			buses = append(buses, bus)
		}
	}

	return timestamp, buses
}

func getAnswerCountPartOne(input string) int {
	timestamp, buses := parse(input)
	min := timestamp
	busID := 0

	// fmt.Printf("\nTimestamp: %d, Buses: %v\n\n", timestamp, buses)
	for _, bus := range buses {
		diff := ((timestamp/bus)*bus + bus) - timestamp
		if diff < min {
			min = diff
			busID = bus
		}
	}

	return min * busID
}

func getAnswerCountPartTwo(input string) int {
	return 1
}

func GetAnswerCountPartOne(input string) int {
	return getAnswerCountPartOne(input)
}

func GetAnswerCountPartTwo(input string) int {
	return getAnswerCountPartTwo(input)
}
