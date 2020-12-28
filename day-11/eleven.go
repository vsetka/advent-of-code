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

func updateSeating(layout [][]string) [][]string {
	updatedSeating := make([][]string, len(layout))
	for i := range layout {
		updatedSeating[i] = make([]string, len(layout[i]))
		copy(updatedSeating[i], layout[i])
	}

	neighborModifiers := [][]int{
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

			for _, mod := range neighborModifiers {
				xMod := x + mod[1]
				yMod := y + mod[0]
				validY := yMod >= 0 && yMod < len(layout)
				validX := validY && xMod >= 0 && xMod < len(layout[y])

				if validX && validY && layout[yMod][xMod] == occupied {
					occupiedCount++
				}
			}

			if layout[y][x] == empty && occupiedCount == 0 {
				updatedSeating[y][x] = occupied
			} else if layout[y][x] == occupied && occupiedCount >= 4 {
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
		layout = updateSeating(layout)
		newCount := countOccupied(layout)
		if newCount == occupiedCount {
			break
		}
		occupiedCount = newCount
	}

	return occupiedCount
}

func getAnswerCountPartTwo(input string) int {
	return 0
}
