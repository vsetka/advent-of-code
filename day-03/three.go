package three

import "strings"

var tree byte = '#'

type MovementPattern struct {
	Right int
	Down  int
}

func countTreesPartOne(input string, movementPattern MovementPattern) int {
	columnIndex := 0
	treesEncountered := 0
	rows := strings.Split(input, "\n")

	for rowIndex := 0; rowIndex < len(rows); rowIndex += movementPattern.Down {
		row := rows[rowIndex]

		if row[columnIndex] == tree {
			treesEncountered++
		}

		columnIndex = (columnIndex + movementPattern.Right) % len(row)
	}

	return treesEncountered
}
