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

	var bags = make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key, value := add(scanner.Text())
		bags[key] = value
	}

	startBag := "shiny gold"
	count := calculate(startBag, getChildren(startBag, bags), bags)
	fmt.Printf("Count: %d\n", count-1)
}

func add(s string) (string, []string) {
	splitted := strings.Split(s, " bags contain ")
	bag := splitted[0]

	if strings.Contains(s, "no other bags") {
		return bag, []string{}
	}

	other := strings.Replace(splitted[1], " bags.", "@", -1)
	other = strings.Replace(other, " bags, ", "@", -1)
	other = strings.Replace(other, " bag.", "@", -1)
	other = strings.Replace(other, " bag, ", "@", -1)

	tokens := strings.Split(other, "@")
	bags := []string{}

	for _, b := range tokens {
		b = strings.TrimSpace(b)
		if len(b) > 0 {
			bags = append(bags, b)
		}
	}

	return bag, bags
}

func getChildren(bag string, bags map[string][]string) []string {
	for k, v := range bags {
		if strings.Contains(bag, k) {
			return v
		}
	}

	return []string{}
}

func calculate(bag string, children []string, bags map[string][]string) int {
	res := 1
	if len(children) == 0 {
		return 1
	}

	for _, c := range children {
		num, _ := strconv.Atoi(c[0:1])
		res += num * calculate(c, getChildren(c, bags), bags)
	}

	return res
}
