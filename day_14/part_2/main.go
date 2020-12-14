package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	mem := make(map[int]int)
	var mask string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mask") {
			mask = getMask(line)
		} else {
			writeToMemory(mem, line, mask)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	fmt.Println(sum)
}

func getMask(s string) string {
	return strings.Split(s, " = ")[1]
}

func writeToMemory(mem map[int]int, s string, mask string) {
	tokens := strings.Split(s, " = ")
	addr := applyMask(getAddress(tokens[0]), mask)
	addresses := mapAddress(addr)
	val, _ := strconv.Atoi(tokens[1])

	for _, a := range convert(addresses) {
		mem[a] = val
	}
}

func applyMask(v int, mask string) string {
	b := fmt.Sprintf("%0*b", len(mask), v)

	var val bytes.Buffer
	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case 'X':
			val.WriteString("X")
		case '1':
			val.WriteString("1")
		case '0':
			val.WriteString(string(b[i]))
		}
	}

	return val.String()
}

func getAddress(s string) int {
	a := strings.ReplaceAll(s, "mem", "")
	a = strings.ReplaceAll(a, "[", "")
	a = strings.ReplaceAll(a, "]", "")
	addr, _ := strconv.Atoi(a)
	return addr
}

func mapAddress(a string) []string {
	addresses := []string{a}
	temp := []string{}
	count := strings.Count(a, "X")
	for i := 0; i < count; i++ {
		for _, addr := range addresses {
			temp = append(temp, strings.Replace(addr, "X", "0", 1))
			temp = append(temp, strings.Replace(addr, "X", "1", 1))
		}
		addresses = temp
		temp = []string{}
	}

	return addresses
}

func convert(s []string) []int {
	result := []int{}
	for _, str := range s {
		num, _ := strconv.ParseInt(str, 2, 64)
		result = append(result, int(num))
	}

	return result
}
