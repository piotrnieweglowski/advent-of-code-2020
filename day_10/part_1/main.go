package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	diff := make(map[int]int)
	diff[1] = 0
	diff[3] = 1
	numbers := []int{0}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)
	for i := 0; i < len(numbers)-1; i++ {
		diff[numbers[i+1]-numbers[i]]++
	}

	fmt.Println(diff)
	fmt.Println(diff[1] * diff[3])
}
