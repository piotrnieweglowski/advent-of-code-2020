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

	acc := 0
	instructions := []instruction{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(tokens[1])
		ins := create(tokens[0], val, &acc)
		instructions = append(instructions, ins)
	}

	stop := false
	ptr := 0
	for !stop {
		ptr, stop = execute(&instructions[ptr], ptr)
	}

	fmt.Println(acc)
}

func execute(ins *instruction, ptr int) (int, bool) {
	if ins.executionCount > 0 {
		return 0, true
	}

	next := ptr + 1
	if ins.code == "acc" {
		*ins.acc += ins.val
	}
	if ins.code == "jmp" {
		next += ins.val - 1
	}

	ins.executionCount++
	return next, false
}

func create(code string, val int, acc *int) instruction {
	return instruction{code, val, 0, acc}
}

type instruction struct {
	code           string
	val            int
	executionCount int
	acc            *int
}
