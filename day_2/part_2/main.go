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
	first  int
	second int
	char   byte
	secret string
}

func create(s string) pass {
	parts := strings.Split(s, ":")
	policy := strings.TrimSpace(parts[0])
	password := strings.TrimSpace(parts[1])
	validation := strings.Split(policy, " ")
	firstAndSecond := strings.TrimSpace(validation[0])
	char := byte(strings.TrimSpace(validation[1])[0])
	firstSecond := strings.Split(firstAndSecond, "-")
	first, _ := strconv.Atoi(firstSecond[0])
	second, _ := strconv.Atoi(firstSecond[1])

	return pass{first, second, char, password}
}

func (p pass) valid() int {
	isFirstOk := (p.secret[p.first-1] == p.char)
	isSecondOk := (p.secret[p.second-1] == p.char)
	if isFirstOk != isSecondOk {
		return 1
	}

	return 0
}
