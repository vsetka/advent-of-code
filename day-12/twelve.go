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

type location struct {
	x int
	y int
}

type waypoint struct {
	position    location
	directionNS int
	directionEW int
}

type ship struct {
	waypoint  location
	position  location
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
func strToIntDirection(val string) int {
	switch val {
	case north:
		return 0
	case east:
		return 1
	case south:
		return 2
	case west:
		return 3
	default:
		fmt.Printf("Invalid value %s. No direction", val)
		return -1
	}
}

func rotatePosition(pos location, direction string, rotations int) location {
	newPos := pos

	// on even rotations, we don't switch axis, we just invert values
	if (rotations % 2) == 0 {
		newPos.x = -newPos.x
		newPos.y = -newPos.y
	} else {
		// on uneven rotations, depending on rotation direction and rotation count
		// we need to not only flip the axis but also change the sign
		if rotations == 1 && direction == left || rotations == 3 && direction == right {
			newPos.x = -pos.y
			newPos.y = pos.x
		} else if rotations == 1 && direction == right || rotations == 3 && direction == left {
			newPos.x = pos.y
			newPos.y = -pos.x
		}
	}

	return newPos
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

func (s *ship) movePartTwo(m move) {
	switch m.direction {
	case north:
		s.waypoint.y += m.value
	case south:
		s.waypoint.y -= m.value
	case east:
		s.waypoint.x += m.value
	case west:
		s.waypoint.x -= m.value
	case left:
		fallthrough
	case right:
		s.waypoint = rotatePosition(s.waypoint, m.direction, (m.value/90)%4)
	case forward:
		s.position.x += s.waypoint.x * m.value
		s.position.y += s.waypoint.y * m.value
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
		position: location{
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
	data := parse(input)
	ship := ship{
		waypoint: location{
			x: 10,
			y: 1,
		},
		position: location{
			x: 0,
			y: 0,
		},
	}

	for _, instruction := range data {
		ship.movePartTwo(instruction)
	}

	return ship.distance()
}
