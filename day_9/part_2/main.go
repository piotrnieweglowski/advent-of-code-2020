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
	found := false
	for i := 0; i < len(numbers); i++ {
		sum := 0
		min := numbers[i]
		max := numbers[i]
		for j := i; j < len(numbers) && !found; j++ {
			sum += numbers[j]
			if numbers[j] < min {
				min = numbers[j]
			}
			if numbers[j] > max {
				max = numbers[j]
			}
			if sum == answer {
				found = true
				fmt.Println("success")
				fmt.Printf("min: %d\n", min)
				fmt.Printf("max: %d\n", max)
				fmt.Printf("together: %d\n", min+max)
			}
		}
	}
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
