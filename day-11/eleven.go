package eleven

import (
	"fmt"
	"strings"
)

const (
	empty    = "L"
	occupied = "#"
)

func parse(input string) [][]string {
	rows := strings.Split(input, "\n")
	data := make([][]string, len(rows))

	for rowIndex, row := range rows {
		for _, cell := range strings.Split(row, "") {
			data[rowIndex] = append(data[rowIndex], cell)
		}
	}

	return data
}

func printSeating(layout [][]string) {
	for y := range layout {
		fmt.Printf("%d:\t%s\n", y, strings.Join(layout[y], ""))
	}

	fmt.Println()
}

func countOccupied(layout [][]string) int {
	occupiedCount := 0

	for _, row := range layout {
		for _, seat := range row {
			if seat == occupied {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func isOccupied(layout [][]string, y int, x int, yModifier int, xModifier int, iterationsLeft int) bool {
	xMod := x + xModifier
	yMod := y + yModifier
	validY := yMod >= 0 && yMod < len(layout)
	validX := validY && xMod >= 0 && xMod < len(layout[y])

	if validX && validY && layout[yMod][xMod] != empty {
		if layout[yMod][xMod] == occupied {
			return true
		}

		iterationsLeft--

		if iterationsLeft > 0 {
			return isOccupied(layout, yMod, xMod, yModifier, xModifier, iterationsLeft)
		}

		return false
	}

	return false
}

func updateSeating(layout [][]string, checkDepth int, maxOccupied int) [][]string {
	updatedSeating := make([][]string, len(layout))
	for i := range layout {
		updatedSeating[i] = make([]string, len(layout[i]))
		copy(updatedSeating[i], layout[i])
	}

	modifiers := [][]int{
		[]int{-1, -1},
		[]int{-1, 0},
		[]int{-1, 1},
		[]int{0, -1},
		[]int{0, 1},
		[]int{1, -1},
		[]int{1, 0},
		[]int{1, 1},
	}

	for y := 0; y < len(layout); y++ {
		for x := 0; x < len(layout[y]); x++ {
			occupiedCount := 0

			for _, mod := range modifiers {
				if isOccupied(layout, y, x, mod[0], mod[1], checkDepth) {
					occupiedCount++
				}
			}

			if layout[y][x] == empty && occupiedCount == 0 {
				updatedSeating[y][x] = occupied
			} else if layout[y][x] == occupied && occupiedCount >= maxOccupied {
				updatedSeating[y][x] = empty
			}
		}
	}

	// printSeating(updatedSeating)
	return updatedSeating
}

func getAnswerCountPartOne(input string) int {
	layout := parse(input)
	occupiedCount := 0

	for {
		layout = updateSeating(layout, 1, 4)
		newCount := countOccupied(layout)
		if newCount == occupiedCount {
			break
		}
		occupiedCount = newCount
	}

	return occupiedCount
}

func getAnswerCountPartTwo(input string) int {
	layout := parse(input)
	occupiedCount := 0

	for {
		layout = updateSeating(layout, len(layout), 5)
		newCount := countOccupied(layout)
		if newCount == occupiedCount {
			break
		}
		occupiedCount = newCount
	}

	return occupiedCount
}
