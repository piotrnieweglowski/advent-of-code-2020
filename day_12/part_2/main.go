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

	wayPoint := &wayPoint{[]int{10, -1}}
	ship := &ship{0, []int{0, 0}}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		move := parseCommand(scanner.Text())
		if move.forward {
			ship.move(*wayPoint, move.steps)
		} else {
			wayPoint.move(move)
		}
	}

	fmt.Println(ship.getManhattanDistance())
}

func parseCommand(s string) move {
	action := s[0]
	value, _ := strconv.Atoi(s[1:len(s)])

	if action == 'F' {
		return move{true, value, 0, 0, 0}
	}
	if action == 'L' || action == 'R' {
		return move{false, 0, 0, value, action}
	}
	return move{false, value, action, 0, 0}
}

type ship struct {
	angle    int
	position []int
}

func (s *ship) move(w wayPoint, steps int) {
	s.position[0] += w.position[0] * steps
	s.position[1] += w.position[1] * steps
}

func (s ship) getManhattanDistance() int {
	x := int(math.Abs(float64(s.position[0])))
	y := int(math.Abs(float64(s.position[1])))
	return x + y
}

type wayPoint struct {
	position []int
}

func (w *wayPoint) move(m move) {
	w.setPosition(m.direction, m.steps)
	w.rotate(m)
}

func (w *wayPoint) setPosition(direction byte, steps int) {
	switch direction {
	case 'N':
		w.position[1] -= steps
	case 'W':
		w.position[0] -= steps
	case 'S':
		w.position[1] += steps
	case 'E':
		w.position[0] += steps
	}
}

func (w *wayPoint) rotate(m move) {
	var deg int
	if m.rotateDirection == 'R' {
		deg = m.rotate
	} else {
		deg = 360 - m.rotate
	}

	x := w.position[0]
	y := w.position[1]
	for deg > 0 {
		w.position[0] = -y
		w.position[1] = x
		x = w.position[0]
		y = w.position[1]
		deg -= 90
	}
}

type move struct {
	forward         bool
	steps           int
	direction       byte
	rotate          int
	rotateDirection byte
}
