package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	var data []string
	rightSteps := []int{1, 3, 5, 7, 1}
	downSteps := []int{1, 1, 1, 1, 2}
	var trees []int
	answer := 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	for i := 0; i < len(rightSteps); i++ {
		posX := 0
		rightStep := rightSteps[i]
		downStep := downSteps[i]
		currentTrees := 0

		row := 0
		for row < len(data) {
			line := data[row]
			currentTrees += isTree(line, posX)
			posX += rightStep
			margin := len(line) - posX
			if margin <= 0 {
				posX = -margin
			}

			row++
			for row % downStep != 0 {
				row++
			}
		}

		trees = append(trees, currentTrees)
		answer *= currentTrees
	}

	fmt.Println(trees)
	fmt.Println(answer)
}

func isTree(s string, pos int) int {
	if s[pos] == '#' {
		return 1
	}
	return 0
}
