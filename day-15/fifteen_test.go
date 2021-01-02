package fifteen

import (
	"testing"
)

var basicInputs = []string{
	`0,3,6`,
	`1,3,2`,
	`2,1,3`,
	`1,2,3`,
	`2,3,1`,
	`3,2,1`,
	`3,1,2`,
}

var advancedInput = `2,0,1,7,4,14,18`

func TestBaseCasePartOne(t *testing.T) {
	var basicOutputs = []int{
		436,
		1,
		10,
		27,
		78,
		438,
		1836,
	}

	for idx := range basicInputs {
		foundSolution := getAnswerCount(basicInputs[idx], 2020)

		if foundSolution != basicOutputs[idx] {
			t.Fatalf("Expected to get %d but got %d\n", basicOutputs[idx], foundSolution)
		}
	}
}

func TestAdvancedCasePartOne(t *testing.T) {
	expectedSolution := 496
	foundSolution := getAnswerCount(advancedInput, 2020)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestBaseCasePartTwo(t *testing.T) {
	var basicOutputs = []int{
		175594,
		2578,
		3544142,
		261214,
		6895259,
		18,
		362,
	}

	for idx := range basicInputs {
		foundSolution := getAnswerCount(basicInputs[idx], 30000000)

		if foundSolution != basicOutputs[idx] {
			t.Fatalf("Expected to get %d but got %d\n", basicOutputs[idx], foundSolution)
		}
	}
}

func TestAdvancedCasePartTwo(t *testing.T) {
	expectedSolution := 883
	foundSolution := getAnswerCount(advancedInput, 30000000)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}
