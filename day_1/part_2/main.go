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

	found := false
	for i := 0; i < len(num) && !found; i++ {
		for j := i + 1; j < len(num) && !found; j++ {
			for k := j + 1; k < len(num) && !found; k++ {
				if num[i]+num[j]+num[k] == year2020 {
					fmt.Printf("First: %d, Second: %d, Third: %d, Result %d * %d * %d = %d\n",
						num[i],
						num[j],
						num[k],
						num[i],
						num[j],
						num[k],
						num[i]*num[j]*num[k])
					found = true
				}
			}
		}
	}
}
