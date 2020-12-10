package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	var answer int
	numbers := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	preambleLength := 25
	for i := preambleLength; i < len(numbers); i++ {
		sum := preamble(numbers[i-preambleLength : i])
		if _, found := sum[numbers[i]]; !found {
			answer = numbers[i]
			break
		}
	}

	fmt.Println(answer)
}

func preamble(n []int) map[int]bool {
	sum := make(map[int]bool)
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if _, ok := sum[n[i]+n[j]]; !ok {
				sum[n[i]+n[j]] = true
			}
		}
	}

	return sum
}
