package thirteen

import (
	"testing"
)

var basicInput = `939
7,13,x,x,59,x,31,19`

var advancedInput = `1003681
23,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,x,x,x,x,x,x,x,x,409,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29`

func TestBaseCasePartOne(t *testing.T) {
	expectedSolution := 295
	foundSolution := getAnswerCountPartOne(basicInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestAdvancedCasePartOne(t *testing.T) {
	expectedSolution := 2045
	foundSolution := getAnswerCountPartOne(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestBaseCasePartTwo(t *testing.T) {
	expectedSolution := 1068781
	foundSolution := getAnswerCountPartTwo(basicInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}

func TestAdvancedCasePartTwo(t *testing.T) {
	expectedSolution := 402251700208309
	foundSolution := getAnswerCountPartTwo(advancedInput)

	if foundSolution != expectedSolution {
		t.Fatalf("Expected to get %d but got %d\n", expectedSolution, foundSolution)
	}
}
