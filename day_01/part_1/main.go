package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const year2020 = 2020
	f, _ := os.Open("data.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var num []int

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		num = append(num, i)
	}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == year2020 {
				fmt.Printf("First: %d, Second: %d, Result %d * %d = %d\n",
					num[i],
					num[j],
					num[i],
					num[j],
					num[i]*num[j])
			}
		}
	}
}
