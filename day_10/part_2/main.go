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

	numbers := []int{0}
	removed := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	ptr := 1
	for ptr < len(numbers)-1 {
		if canRemove(numbers[ptr-1 : ptr+2]) {
			removed = append(removed, numbers[ptr])
		}
		ptr++
	}

	fmt.Println(removed)

	groups := group(removed)
	combinationCounts := getCombinationCount(groups)
	ans := 1
	for _, v := range combinationCounts {
		ans *= v
	}

	fmt.Println(ans)
}

func canRemove(tree []int) bool {
	if len(tree) != 3 {
		panic("Slice must contain exact 3 numbers")
	}

	diff := tree[2] - tree[0]
	if diff >= 1 && diff <= 3 {
		return true
	}

	return false
}

func group(removed []int) (groups []int) {
	g := 1
	for i := 1; i < len(removed); i++ {
		if removed[i]-removed[i-1] != 1 {
			groups = append(groups, g)
			g = 0
		}
		g++
		if i+1 == len(removed) {
			groups = append(groups, g)
		}
	}

	return groups
}

func getCombinationCount(groups []int) (mapped []int) {
	for _, v := range groups {
		switch v {
		case 1:
			mapped = append(mapped, 2)
		case 2:
			mapped = append(mapped, 4)
		case 3:
			mapped = append(mapped, 7)
		}
	}

	return mapped
}
