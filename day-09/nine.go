package nine

import (
	"fmt"
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

func toMap(input []int) map[int]int {
	var mapped = make(map[int]int, len(input))

	for _, item := range input {
		mapped[item] = item
	}

	return mapped
}

func getAnswerCountPartOne(input string, preambleSize int) int {
	data := parse(input)
	preamble := toMap(data[:preambleSize])
	rest := data[preambleSize:]
	answer := 0
	valid := false

	for idx, item := range rest {
		valid = false

		for preambleItem := range preamble {
			if _, ok := preamble[item-preambleItem]; ok {
				valid = true
				preamble = toMap(data[idx+1 : preambleSize+idx+1])
				break
			}
		}

		if !valid {
			answer = item
			break
		}
	}

	return answer
}

func getAnswerCountPartTwo(input string) int {
	return 0
}
