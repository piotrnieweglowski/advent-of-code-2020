package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pos := 0
	rightStep := 3
	trees := 0
	for scanner.Scan() {
		line := scanner.Text()
		trees += isTree(line, pos)
		pos += rightStep

		margin := len(line) - pos
		if margin <= 0 {
			pos = -margin
		}
	}

	fmt.Printf("Trees count: %d\n", trees)
}

func isTree(s string, pos int) int {
	if s[pos] == '#' {
		return 1
	}
	return 0
}
