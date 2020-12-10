package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var seats = []int{}
	myID := 0

	for scanner.Scan() {
		seats = append(seats, id(scanner.Text()))
	}

	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i]+1 != seats[i+1] {
			myID = seats[i] + 1
		}
	}

	fmt.Println(myID)
}

func id(s string) int {
	id := partition(s[0:7], 'F') * 8
	id += partition(s[7:10], 'L')
	return id
}

func partition(s string, lower byte) int {
	min := 0.0
	max := math.Pow(2, float64(len(s))) - 1

	for i := 0; i < len(s); i++ {
		avg := (min + max) / 2
		if s[i] == lower {
			max = math.Floor(avg)
		} else {
			min = math.Ceil(avg)
		}
	}

	if min != max {
		panic("Critical error")
	}

	return int(min)
}
