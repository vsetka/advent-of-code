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
	baseJoltage := 0 // outlet is 0, so we start with 0

	sort.Slice(data, func(a int, b int) bool {
		return data[a] < data[b]
	})

	for _, adapterJoltage := range data {
		differenceCounts[adapterJoltage-baseJoltage]++
		baseJoltage = adapterJoltage
	}

	differenceCounts[3]++ // accounts for our device

	return differenceCounts[1] * differenceCounts[3]
}

func getAnswerCountPartTwo(input string) int {
	data := parse(input)
	combinationCounts := map[int]int{0: 1}

	sort.Slice(data, func(a int, b int) bool {
		return data[a] < data[b]
	})

	for _, adapterJoltage := range data {
		for joltageDifference := 1; joltageDifference <= 3; joltageDifference++ {
			if val, ok := combinationCounts[(adapterJoltage - joltageDifference)]; ok {
				combinationCounts[adapterJoltage] += val
			}
		}
	}

	return combinationCounts[data[(len(data)-1)]]
}
