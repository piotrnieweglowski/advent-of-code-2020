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
	myTicket := extractTickets(lines[index])

	index += 3
	allPositions := []int{}
	nearbyTickets := [][]int{}
	for i := index; i < len(lines); i++ {
		partial := extractTickets(lines[i])
		nearbyTickets = append(nearbyTickets, partial)
		for _, p := range partial {
			allPositions = append(allPositions, p)
		}
	}

	invalid := []int{}
	for _, t := range allPositions {
		if !contains(t, rules) {
			invalid = append(invalid, t)
		}
	}

	validTickets := [][]int{}
	for _, t := range nearbyTickets {
		if !containsInvalid(t, invalid) {
			validTickets = append(validTickets, t)
		}
	}

	ticketPositions := make(map[string]int)
	for i := 0; i < len(myTicket); i++ {
		positions := getTicketsPosition(i, validTickets)
		for j := 0; j < len(rules); j += 2 {
			if containsTicketPosition(rules[j].description, ticketPositions) {
				continue
			}
			if satisfiesAll(rules[j:j+2], positions) {
				ticketPositions[rules[j].description] = i
				break
			}
		}
	}

	fmt.Println(myTicket)
	answer := 1
	for t, i := range ticketPositions {
		if strings.Contains(t, "departure") {
			answer *= myTicket[i]
		}
	}

	fmt.Println(answer)
}

func extractTickets(s string) (ticket []int) {
	for _, num := range strings.Split(s, ",") {
		n, _ := strconv.Atoi(num)
		ticket = append(ticket, n)
	}
	return ticket
}

type rule struct {
	description string
	min         int
	max         int
}

func create(s string) (rules []rule) {
	tokens := strings.Split(s, ": ")
	ranges := strings.Split(tokens[1], " or ")
	for _, r := range ranges {
		minMax := strings.Split(r, "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		rules = append(rules, rule{tokens[0], min, max})
	}
	return rules
}

func (r rule) contains(num int) bool {
	return r.min <= num && r.max >= num
}

func satisfiesAll(rules []rule, values []int) bool {
	if rules[0].description != rules[1].description || len(rules) != 2 {
		panic("Critical error")
	}

	for _, v := range values {
		if !rules[0].contains(v) && !rules[1].contains(v) {
			return false
		}
	}
	return true
}

func contains(num int, rules []rule) bool {
	for _, r := range rules {
		if r.contains(num) {
			return true
		}
	}
	return false
}

func containsInvalid(ticket []int, invalid []int) bool {
	for _, t := range ticket {
		for _, i := range invalid {
			if t == i {
				return true
			}
		}
	}

	return false
}

func getTicketsPosition(position int, tickets [][]int) []int {
	result := []int{}
	for i := 0; i < len(tickets); i++ {
		result = append(result, tickets[i][position])
	}
	return result
}

func containsTicketPosition(position string, positions map[string]int) bool {
	for p := range positions {
		if position == p {
			return true
		}
	}
	return false
}
