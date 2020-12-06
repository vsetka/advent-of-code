package six

import (
	"regexp"
	"strings"
)

func getAnswerCountPartOne(input string) int {
	groupDelimiter := regexp.MustCompile(`\n{2,}`)
	uniqueAnswerCount := 0

	for _, group := range groupDelimiter.Split(input, -1) {
		answers := make(map[rune]int)
		for _, row := range strings.Split(group, "\n") {
			for _, answer := range row {
				answers[answer] = 1
			}
		}
		uniqueAnswerCount += len(answers)
	}

	return uniqueAnswerCount
}

func getAnswerCountPartTwo(input string) int {
	groupDelimiter := regexp.MustCompile(`\n{2,}`)
	uniqueAnswerCount := 0

	for _, group := range groupDelimiter.Split(input, -1) {
		answers := make(map[rune]int)
		rows := strings.Split(group, "\n")
		for _, row := range rows {
			for _, answer := range row {
				answers[answer]++
				if answers[answer] == len(rows) {
					uniqueAnswerCount++
				}
			}
		}
	}

	return uniqueAnswerCount
}
