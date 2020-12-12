package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	ship := &ship{'E', 0, []int{0, 0}}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ship.move(scanner.Text())
	}

	fmt.Println(ship.getManhattanDistance())
}

type ship struct {
	direction byte
	angle     int
	position  []int
}

func (s *ship) move(command string) {
	move := parseCommand(command)
	if move.direction == 0 && move.rotateDirection == 0 {
		s.setPosition(s.direction, move.steps)
	}

	s.setPosition(move.direction, move.steps)
	s.calculateAngle(move)
}

func (s *ship) setPosition(direction byte, steps int) {
	switch direction {
	case 'N':
		s.position[1] -= steps
	case 'W':
		s.position[0] -= steps
	case 'S':
		s.position[1] += steps
	case 'E':
		s.position[0] += steps
	}
}

func (s *ship) calculateAngle(move move) {
	if move.rotateDirection == 'L' {
		s.angle += (360 - move.rotate)
	} else if move.rotateDirection == 'R' {
		s.angle += move.rotate
	}

	s.angle = s.angle % 360
	if s.angle >= 0 && s.angle < 90 {
		s.direction = 'E'
	} else if s.angle >= 90 && s.angle < 180 {
		s.direction = 'S'
	} else if s.angle >= 180 && s.angle < 270 {
		s.direction = 'W'
	} else {
		s.direction = 'N'
	}
}

func (s ship) getManhattanDistance() int {
	x := int(math.Abs(float64(s.position[0])))
	y := int(math.Abs(float64(s.position[1])))
	return x + y
}

func parseCommand(s string) move {
	action := s[0]
	value, _ := strconv.Atoi(s[1:len(s)])

	if action == 'F' {
		return move{value, 0, 0, 0}
	}
	if action == 'L' || action == 'R' {
		return move{0, 0, value, action}
	}
	return move{value, action, 0, 0}
}

type move struct {
	steps           int
	direction       byte
	rotate          int
	rotateDirection byte
}
