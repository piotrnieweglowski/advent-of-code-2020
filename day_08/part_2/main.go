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

	ptr := 0
	replaced := make(map[int]bool)
	for ptr < len(instructions) {
		reset(instructions)
		stop := false
		acc = 0
		ptr = 0
		replacedCount := len(replaced)
		for !stop && ptr < len(instructions) {
			canReplace := replacedCount == len(replaced)
			ptr, stop = execute(&instructions[ptr], ptr, replaced, canReplace)
		}
	}

	fmt.Println(acc)
}

func execute(ins *instruction, ptr int, replaced map[int]bool, canReplace bool) (int, bool) {
	if ins.executionCount > 0 {
		return 0, true
	}

	next := ptr + 1
	_, ok := replaced[ptr]
	if ins.code == "acc" {
		*ins.acc += ins.val
	}
	if ins.code == "jmp" {
		if !ok && canReplace {
			replaced[ptr] = true
		} else {
			next += ins.val - 1
		}
	}
	if ins.code == "nop" && !ok && canReplace {
		replaced[ptr] = true
		next += ins.val - 1
	}

	ins.executionCount++
	return next, false
}

func reset(instructions []instruction) {
	for i := 0; i < len(instructions); i++ {
		resetExecutionCount(&instructions[i])
	}
}

func resetExecutionCount(ins *instruction) {
	ins.executionCount = 0
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
