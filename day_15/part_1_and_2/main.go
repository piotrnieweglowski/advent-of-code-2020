package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	numbers := []int{}
	numbersSpoken := make(map[int]int)
	last := 0
	next := 0

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ",")
		for i, s := range data {
			n, _ := strconv.Atoi(s)
			if i != len(data)-1 {
				numbers = append(numbers, n)
				numbersSpoken[n] = i + 1
			}
			last = n
		}
	}

	// part 1: searched 2020
	// part 2: searched 30000000
	searched := 30000000
	for i := len(numbers) + 1; i <= searched; i++ {
		next = generateNext(last, i, numbersSpoken)
		numbers = append(numbers, last)
		numbersSpoken[last] = i
		last = next
	}

	fmt.Println(numbers[searched-1])
}

func generateNext(num int, turn int, numbersSpoken map[int]int) int {
	if lastTurn, ok := numbersSpoken[num]; ok {
		return turn - lastTurn
	}

	return 0
}
