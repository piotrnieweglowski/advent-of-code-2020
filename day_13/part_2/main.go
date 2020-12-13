package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	scanner.Scan()
	raw := []int{}
	busses := make(map[int]int)
	sorted := []int{}
	max := 0
	index := 0
	for _, v := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(v)
		raw = append(raw, num)
	}

	for i, v := range raw {
		if i+v < len(raw) && raw[i+v] != 0 {
			raw[i] = 0
			raw[i+v] *= v
		}
	}
	fmt.Println(raw)

	for i, v := range raw {
		if v > max {
			max = v
			index = i
		}
		if v != 0 {
			sorted = append(sorted, -v)
			busses[v] = i
		}
	}

	sort.Ints(sorted)
	for i := 0; i < len(sorted); i++ {
		sorted[i] = -sorted[i]
	}
	fmt.Println(sorted)

	timestamp := uint64(max)
	stop := false
	for !stop {
		timestamp += uint64(max)
		stop = true
		for _, key := range sorted {
			i := busses[key]
			toCheck := timestamp + uint64(i-index)
			if toCheck%uint64(key) != 0 {
				stop = false
				break
			}
		}
	}
	fmt.Println(timestamp - uint64(index))
}
