package ten

import (
	"fmt"
	"sort"
	"strings"
)

func parse(input string) []int {
	var data []int
	var value int

	for _, row := range strings.Split(input, "\n") {
		if _, err := fmt.Sscanf(row, "%d", &value); err == nil {
			data = append(data, value)
		}
	}

	return data
}

func getAnswerCountPartOne(input string) int {
	data := parse(input)
	differenceCounts := map[int]int{}
	base := 0 // outlet is 0, so we start with 0

	sort.Slice(data, func(a int, b int) bool {
		return data[a] < data[b]
	})

	for _, item := range data {
		differenceCounts[item-base]++
		base = item
	}

	differenceCounts[3]++ // accounts for our device

	return differenceCounts[1] * differenceCounts[3]
}

func getAnswerCountPartTwo(input string, target int) int {
	// data := parse(input)
	answer := 0

	return answer
}
