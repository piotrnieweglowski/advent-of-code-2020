package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	arrival, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	min, difference := math.MaxInt64, math.MaxInt64
	bus := 0
	scheduleData := strings.Replace(scanner.Text(), "x,", "", -1)
	for _, v := range strings.Split(scheduleData, ",") {
		t, _ := strconv.Atoi(v)
		timestamp, diff := firstTimestampAfter(arrival, t)
		if timestamp < min {
			min = timestamp
			difference = diff
			bus = t
		}
	}

	fmt.Println(bus * difference)
}

func firstTimestampAfter(base int, period int) (timestamp int, difference int) {
	result := 0
	for result < base {
		result += period
	}

	timestamp = result
	difference = result - base
	return timestamp, difference
}
