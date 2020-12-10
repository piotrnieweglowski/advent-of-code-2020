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

	validCount := 0
	for scanner.Scan() {
		p := create(scanner.Text())
		validCount += p.valid()
	}

	fmt.Printf("Valid count: %d\n", validCount)
}

type pass struct {
	min    int
	max    int
	char   string
	secret string
}

func create(s string) pass {
	parts := strings.Split(s, ":")
	policy := strings.TrimSpace(parts[0])
	password := strings.TrimSpace(parts[1])
	validation := strings.Split(policy, " ")
	minWithMax := strings.TrimSpace(validation[0])
	char := strings.TrimSpace(validation[1])
	minMax := strings.Split(minWithMax, "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return pass{min, max, char, password}
}

func (p pass) valid() int {
	c := strings.Count(p.secret, p.char)
	if c >= p.min && c <= p.max {
		return 1
	}

	return 0
}
