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
		} else {
			buses = append(buses, 0)
		}
	}

	return timestamp, buses
}

func getAnswerCountPartOne(input string) int {
	timestamp, buses := parse(input)
	min := timestamp
	busID := 0

	for _, bus := range buses {
		if bus <= 0 {
			continue
		}

		diff := ((timestamp/bus)*bus + bus) - timestamp

		if diff < min {
			min = diff
			busID = bus
		}
	}

	return min * busID
}

func getAnswerCountPartTwo(input string) int {
	_, buses := parse(input)
	timestamp := 1
	step := 1

	for idx, bus := range buses {
		if bus <= 0 {
			continue
		}

		// advance until we find a timestamp which gets divided
		// by the current bus without remainder
		for ((timestamp + idx) % bus) != 0 {
			timestamp += step
		}

		// multiply the step by the bus as we from now on
		// only want to increase the timestamp in lock-step for
		// all buses that have been found so far (zero remainder division)
		step *= bus
	}

	return timestamp
}
