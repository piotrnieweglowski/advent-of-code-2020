package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	group := ""
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 0 {
			group += line
		} else {
			count += answers(group)
			group = ""
		}
	}

	if len(line) > 0 {
		count += answers(group)
	}

	fmt.Printf("Count: %d\n", count)
}

func answers(s string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if !ok {
			m[s[i]] = true
		}
	}

	return len(m)
}
