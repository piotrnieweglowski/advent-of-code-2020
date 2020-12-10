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
	group := []string{}
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) > 0 {
			group = append(group, line)
		} else {
			count += answers(group)
			group = []string{}
		}
	}

	if len(line) > 0 {
		count += answers(group)
	}

	fmt.Printf("Count: %d\n", count)
}

func answers(s []string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			_, ok := m[s[i][j]]
			if !ok {
				m[s[i][j]] = 1
			} else {
				m[s[i][j]]++
			}
		}
	}

	all := 0
	for _, v := range m {
		if v == len(s) {
			all++
		}
	}

	return all
}
