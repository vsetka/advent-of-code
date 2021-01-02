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

var basicOutputs = []int{
	436,
	1,
	10,
	27,
	78,
	438,
	1836,
}

var advancedInput = `2,0,1,7,4,14,18`

func TestBaseCasePartOne(t *testing.T) {
	for idx := range basicInputs {
		foundSolution := getAnswerCountPartOne(basicInputs[idx])

		if foundSolution != basicOutputs[idx] {
			t.Fatalf("Expected to get %d but got %d\n", basicOutputs[idx], foundSolution)
		}
	}
}

func TestAdvancedCasePartOne(t *testing.T) {
	expectedSolution := 496
	foundSolution := getAnswerCountPartOne(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestBaseCasePartTwo(t *testing.T) {
	for idx := range basicInputs {
		foundSolution := getAnswerCountPartTwo(basicInputs[idx])

		if foundSolution != basicOutputs[idx] {
			t.Fatalf("Expected to get %d but got %d\n", basicOutputs[idx], foundSolution)
		}
	}
}

func TestAdvancedCasePartTwo(t *testing.T) {
	expectedSolution := 1
	foundSolution := getAnswerCountPartTwo(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}
