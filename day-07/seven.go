package seven

import (
	"regexp"
	"strconv"
	"strings"
)

type capacity struct {
	quantity int
	bag      string
}

func getCapacityMap(input string) map[string][]capacity {
	rootMatcher := regexp.MustCompile(`(?P<rootBag>\w+\s+\w+)\s+bags contain\s+(?P<rest>.+)\.`)
	childMatcher := regexp.MustCompile(`(\d+)\s+(\w+\s+\w+)`)
	capacities := make(map[string][]capacity)

	for _, row := range strings.Split(input, "\n") {
		fieldMatches := rootMatcher.FindAllStringSubmatch(row, -1)

		for _, match := range fieldMatches {
			rootBag := match[1]
			rest := match[2]
			// fmt.Printf("rootBag: %s\n", rootBag)

			if rest != "no other bags" {
				restMatches := childMatcher.FindAllStringSubmatch(rest, -1)
				for _, childMatch := range restMatches {
					quantity, _ := strconv.Atoi(childMatch[1])
					childBag := childMatch[2]
					capacities[rootBag] = append(capacities[rootBag], capacity{quantity: quantity, bag: childBag})
					// fmt.Printf("childMatch: %d %s\n", quantity, childBag)
				}
			}
		}
	}

	return capacities
}

func hasBag(capacities []capacity, bag string) bool {
	for _, c := range capacities {
		if c.bag == bag {
			return true
		}
	}

	return false
}

func contains(items []string, target string) bool {
	for _, i := range items {
		if i == target {
			return true
		}
	}

	return false
}

func distinct(items []string) []string {
	d := []string{}

	for _, i := range items {
		if !contains(d, i) {
			d = append(d, i)
		}
	}

	return d
}

func getHolders(bag string, capacitiesMap map[string][]capacity) []string {
	additionalCandidates := []string{}

	for parent, children := range capacitiesMap {
		if hasBag(children, bag) {
			additionalCandidates = append(additionalCandidates, parent)
			additionalCandidates = append(additionalCandidates, getHolders(parent, capacitiesMap)...)
		}
	}

	return additionalCandidates
}

// GetAnswerCountPartOne asd
func GetAnswerCountPartOne(input string) int {
	capacities := getCapacityMap(input)
	holders := []string{}

	candidates := make([]string, 0, len(capacities))
	for k := range capacities {
		candidates = append(candidates, k)
	}

	holders = distinct(getHolders("shiny gold", capacities))

	return len(holders)
}
