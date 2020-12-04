package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	passport := ""
	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 0 {
			passport += line
		} else {
			count += valid(passport)
			passport = ""
		}
	}

	if len(line) > 0 {
		count += valid(passport)
	}

	fmt.Printf("Valid passport count: %d\n", count)
}

func valid(s string) int {
	// byr (Birth Year)
	// iyr (Issue Year)
	// eyr (Expiration Year)
	// hgt (Height)
	// hcl (Hair Color)
	// ecl (Eye Color)
	// pid (Passport ID)
	// cid (Country ID) - optional
	mandatory := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

	for _, code := range mandatory {
		if !strings.Contains(s, code) {
			return 0
		}
	}
	return 1
}
