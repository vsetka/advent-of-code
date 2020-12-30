package twelve

import (
	"fmt"
	"strings"
)

const (
	north   = "N"
	south   = "S"
	east    = "E"
	west    = "W"
	left    = "L"
	right   = "R"
	forward = "F"
)

type move struct {
	direction string
	value     int
}

type position struct {
	x int
	y int
}

type ship struct {
	position  position
	direction int
}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

// if only there were enums...
func intToStrDirection(val int) string {
	switch val {
	case 0:
		return north
	case 1:
		return east
	case 2:
		return south
	case 3:
		return west
	default:
		fmt.Printf("Invalid value %d. No direction", val)
		return ""
	}
}

func (s *ship) movePartOne(m move) {
	switch m.direction {
	case north:
		s.position.y += m.value
	case south:
		s.position.y -= m.value
	case east:
		s.position.x += m.value
	case west:
		s.position.x -= m.value
	case left:
		s.direction = (s.direction - (m.value / 90)) % 4
		if s.direction < 0 {
			s.direction += 4
		}
	case right:
		s.direction = (s.direction + (m.value / 90)) % 4
	case forward:
		s.movePartOne(move{
			direction: intToStrDirection(s.direction),
			value:     m.value,
		})
	}
}

func (s ship) distance() int {
	return abs(s.position.x) + abs(s.position.y)
}

func parse(input string) []move {
	var data []move
	var direction string
	var value int

	for _, row := range strings.Split(input, "\n") {
		if _, err := fmt.Sscanf(row, "%1s%d", &direction, &value); err == nil {
			data = append(data, move{
				direction: direction,
				value:     value,
			})
		}
	}

	return data
}

func getAnswerCountPartOne(input string) int {
	data := parse(input)
	ship := ship{
		position: position{
			x: 0,
			y: 0,
		},
		// north: 0, east: 1, south: 2, west: 3
		// initially facing east
		direction: 1,
	}

	for _, instruction := range data {
		ship.movePartOne(instruction)
	}

	return ship.distance()
}

func getAnswerCountPartTwo(input string) int {
	return 0
}
