package one

import (
	"errors"
	"sort"
)

func productOfTwo(input []int, sumTarget int) (int, error) {
	candidates := map[int]bool{}

	// Go over all the items
	for _, item := range input {
		// Check the hashmap for an "inverse" of our target and current item.
		// If it's there, we have seen that number before and it sums to the target
		// with our current item. We found the solution!
		if _, ok := candidates[sumTarget-item]; ok == true {
			return item * (sumTarget - item), nil
		}
		// Add the current number to the hashmap
		candidates[item] = true
	}

	return 0, errors.New("No sum found for provided target")
}

func productOfThree(input []int, sumTarget int) (int, error) {
	// Sort the array in ascending order for the algorithm below to work
	sort.Ints(input)

	// Go over the entire array - 2 last members to account for next and last array indices
	for first := range input[:len(input)-2] {
		// Set the next and last array pointers
		next := first + 1
		last := len(input) - 1

		// Move the next and last array indices towards each other...
		for {
			// ... and break out of the loop if they meet
			if next >= last {
				break
			}

			if input[first]+input[next]+input[last] == sumTarget { // Bingo!
				return input[first] * input[next] * input[last], nil
			} else if input[first]+input[next]+input[last] < sumTarget { // Sum under target, keep moving towards the bigger array numbers
				next++
			} else if input[first]+input[next]+input[last] > sumTarget { // Sum is over the target, move last index pointer towards the smaller numbers
				last--
			}
		}
	}

	return 0, errors.New("No sum found for provided target")
}
