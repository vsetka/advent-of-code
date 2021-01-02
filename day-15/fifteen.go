package fifteen

import (
	"strconv"
	"strings"
)

func getAnswerCountPartOne(input string) int {
	inputNumbers := strings.Split(input, ",")
	turn := 0
	spokenNumberRounds := map[int][]int{}
	spokenNumbers := []int{}

	for idx, input := range inputNumbers {
		turn = idx + 1
		inputNumber, _ := strconv.Atoi(input)

		spokenNumbers = append(spokenNumbers, inputNumber)
		spokenNumberRounds[inputNumber] = append(spokenNumberRounds[inputNumber], turn)
	}

	for {
		turn++
		lastSpoken := spokenNumbers[len(spokenNumbers)-1]
		spokenTurns := spokenNumberRounds[lastSpoken]
		spokenTimes := len(spokenTurns)

		if spokenTimes == 1 {
			spokenNumbers = append(spokenNumbers, 0)
			spokenNumberRounds[0] = append(spokenNumberRounds[0], turn)
		} else {
			number := spokenTurns[spokenTimes-1] - spokenTurns[spokenTimes-2]
			spokenNumbers = append(spokenNumbers, number)
			spokenNumberRounds[number] = append(spokenNumberRounds[number], turn)
		}
		if len(spokenNumbers) == 2020 {
			return spokenNumbers[2019]
		}
	}
}

func getAnswerCountPartTwo(input string) int {
	return 0
}

func GetAnswerCountPartOne(input string) int {
	return getAnswerCountPartOne(input)
}

func GetAnswerCountPartTwo(input string) int {
	return getAnswerCountPartTwo(input)
}
