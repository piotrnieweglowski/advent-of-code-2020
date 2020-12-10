package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	max := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		id := id(scanner.Text())
		if id > max {
			max = id
		}
	}

	fmt.Printf("Max id: %d\n", max)
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
