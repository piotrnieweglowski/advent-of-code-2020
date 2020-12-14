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
	addr := getAddress(tokens[0])
	val, _ := strconv.Atoi(tokens[1])
	val = applyMask(val, mask)

	mem[addr] = val
}

func applyMask(v int, mask string) int {
	b := fmt.Sprintf("%0*b", len(mask), v)

	var val bytes.Buffer
	for i := 0; i < len(mask); i++ {
		if mask[i] != 'X' {
			val.WriteString(string(mask[i]))
		} else {
			val.WriteString(string(b[i]))
		}
	}

	result, _ := strconv.ParseInt(val.String(), 2, 64)
	return int(result)
}

func getAddress(s string) int {
	a := strings.ReplaceAll(s, "mem", "")
	a = strings.ReplaceAll(a, "[", "")
	a = strings.ReplaceAll(a, "]", "")
	addr, _ := strconv.Atoi(a)
	return addr
}
