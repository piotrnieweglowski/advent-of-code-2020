package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := []string{}
	rules := []rule{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	index := 0
	for index < len(lines) && len(lines[index]) > 0 {
		for _, r := range create(lines[index]) {
			rules = append(rules, r)
		}
		index++
	}

	index += 2

	index += 3
	nearbyTickets := []int{}
	for i := index; i < len(lines); i++ {
		partial := extractTickets(lines[i])
		for _, p := range partial {
			nearbyTickets = append(nearbyTickets, p)
		}
	}

	invalid := []int{}
	for _, t := range nearbyTickets {
		if !contains(t, rules) {
			invalid = append(invalid, t)
		}
	}

	sum := 0
	for _, t := range invalid {
		sum += t
	}

	fmt.Println(sum)
}

func extractTickets(s string) (ticket []int) {
	for _, num := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(num)
		ticket = append(ticket, n)
	}
	return ticket
}

type rule struct {
	min int
	max int
}

func create(s string) (rules []rule) {
	ranges := strings.Split(strings.Split(s, ": ")[1], " or ")
	for _, r := range ranges {
		minMax := strings.Split(r, "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		rules = append(rules, rule{min, max})
	}
	return rules
}

func (r rule) contains(num int) bool {
	return r.min <= num && r.max >= num
}

func contains(num int, rules []rule) bool {
	for _, r := range rules {
		if r.contains(num) {
			return true
		}
	}
	return false
}
