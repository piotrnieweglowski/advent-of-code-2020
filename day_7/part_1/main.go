package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	var bags = make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if key, value, ok := add(scanner.Text()); ok {
			bags[key] = value
		}
	}

	startBag := "shiny gold"
	delete(bags, startBag)

	colors := make(map[string]bool)
	colors[startBag] = true

	added := true
	for added && len(bags) > 0 {
		added = false
		for color := range colors {
			for key, value := range bags {
				_, ok := colors[key]
				if !ok && contain(value, color) {
					colors[key] = true
					added = true
				}
			}
		}

		for color := range colors {
			delete(bags, color)
		}
	}

	delete(colors, startBag)
	fmt.Printf("Count: %d\n", len(colors))
}

func add(s string) (string, []string, bool) {
	if strings.Contains(s, "no other bags") {
		return "", []string{}, false
	}

	splitted := strings.Split(s, " bags contain ")
	bag := splitted[0]

	other := strings.Replace(splitted[1], " bags.", "@", -1)
	other = strings.Replace(other, " bags, ", "@", -1)
	other = strings.Replace(other, " bag.", "@", -1)
	other = strings.Replace(other, " bag, ", "@", -1)

	tokens := strings.Split(other, "@")
	bags := []string{}

	for _, b := range tokens {
		reg, err := regexp.Compile("[^a-zA-Z ]+")
		if err != nil {
			log.Fatal(err)
		}
		b = reg.ReplaceAllString(b, "")
		b = strings.TrimSpace(b)
		if len(b) > 0 {
			bags = append(bags, b)
		}
	}

	return bag, bags, true
}

func contain(container []string, color string) bool {
	for _, c := range container {
		if c == color {
			return true
		}
	}

	return false
}
