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

	count := 0
	scanner := bufio.NewScanner(f)
	passport := ""
	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 0 {
			passport += " " + line
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
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.
	tokens := strings.Split(s, " ")
	if byr(tokens) &&
		iyr(tokens) &&
		eyr(tokens) &&
		hgt(tokens) &&
		hcl(tokens) &&
		ecl(tokens) &&
		pid(tokens) {
		return 1
	}

	return 0
}

func byr(tokens []string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	token := getToken(tokens, "byr")
	year, err := strconv.Atoi(token)
	return len(token) > 0 && err == nil && year >= 1920 && year <= 2002
}

func iyr(tokens []string) bool {
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	token := getToken(tokens, "iyr")
	year, err := strconv.Atoi(token)
	return len(token) > 0 && err == nil && year >= 2010 && year <= 2020
}

func eyr(tokens []string) bool {
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	token := getToken(tokens, "eyr")
	year, err := strconv.Atoi(token)
	return len(token) > 0 && err == nil && year >= 2020 && year <= 2030
}

func hgt(tokens []string) bool {
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	token := getToken(tokens, "hgt")
	if len(token) <= 0 {
		return false
	}

	if strings.Contains(token, "in") {
		t := strings.Trim(token, "in")
		in, err := strconv.Atoi(t)
		return err == nil && in >= 59 && in <= 76
	}

	if strings.Contains(token, "cm") {
		t := strings.Trim(token, "cm")
		cm, err := strconv.Atoi(t)
		return err == nil && cm >= 150 && cm <= 193
	}

	return false
}

func hcl(tokens []string) bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	token := getToken(tokens, "hcl")
	if len(token) <= 0 || len(token) > 7 {
		return false
	}
	if token[0] != '#' {
		return false
	}

	for i := 1; i < len(token); i++ {
		if (token[i] < '0' || token[i] > '9') && (token[i] < 'a' || token[i] > 'f') {
			return false
		}
	}

	return true
}

func ecl(tokens []string) bool {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	token := getToken(tokens, "ecl")
	return token == "amb" ||
		token == "blu" ||
		token == "brn" ||
		token == "gry" ||
		token == "grn" ||
		token == "hzl" ||
		token == "oth"
}

func pid(tokens []string) bool {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	token := getToken(tokens, "pid")
	if len(token) <= 0 || len(token) != 9 {
		return false
	}

	_, err := strconv.Atoi(token)
	if err == nil {
		return true
	}

	return false
}

func getToken(tokens []string, name string) string {
	for _, t := range tokens {
		if strings.Contains(t, name+":") {
			token := strings.Split(t, ":")[1]
			token = strings.TrimSpace(token)
			return token
		}
	}

	return ""
}
