package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("data.txt")
	defer f.Close()

	board := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte{}
		for i := 0; i < len(line); i++ {
			row = append(row, line[i])
		}
		board = append(board, row)
	}

	allPlacesTaken := false
	for !allPlacesTaken {
		board = takeSeats(board)
		board = releaseSeats(board)
		allPlacesTaken = !anyPlaceToTake(board)
	}

	fmt.Println(calculateTakenSeats(board))
}

func takeSeats(board [][]byte) (newBoard [][]byte) {
	for r := 0; r < len(board); r++ {
		row := []byte{}
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] == 'L' && notVisibleDirections(board, r, c) == 0 {
				row = append(row, '#')
			} else {
				row = append(row, board[r][c])
			}
		}
		newBoard = append(newBoard, row)
	}

	return newBoard
}

func releaseSeats(board [][]byte) (newBoard [][]byte) {
	for r := 0; r < len(board); r++ {
		row := []byte{}
		for c := 0; c < len(board[r]); c++ {
			if board[r][c] == '#' && notVisibleDirections(board, r, c) >= 5 {
				row = append(row, 'L')
			} else {
				row = append(row, board[r][c])
			}
		}
		newBoard = append(newBoard, row)
	}

	return newBoard
}

func notVisibleDirections(board [][]byte, row int, column int) int {
	n := visibility(board, row-1, column, []int{-1, 0})
	ne := visibility(board, row-1, column+1, []int{-1, 1})
	e := visibility(board, row, column+1, []int{0, 1})
	se := visibility(board, row+1, column+1, []int{1, 1})
	s := visibility(board, row+1, column, []int{1, 0})
	sw := visibility(board, row+1, column-1, []int{1, -1})
	w := visibility(board, row, column-1, []int{0, -1})
	nw := visibility(board, row-1, column-1, []int{-1, -1})

	return n + ne + e + se + s + sw + w + nw
}

func visibility(board [][]byte, row int, column int, step []int) int {
	rows := len(board)
	columns := len(board[0])

	if row < 0 || row >= rows {
		return 0
	}

	if column < 0 || column >= columns {
		return 0
	}

	if board[row][column] == 'L' {
		return 0
	}

	if board[row][column] == '#' {
		return 1
	}

	return visibility(board, row+step[0], column+step[1], step)
}

func anyPlaceToTake(board [][]byte) bool {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 'L' && notVisibleDirections(board, r, c) == 0 {
				return true
			}
		}
	}

	return false
}

func calculateTakenSeats(board [][]byte) int {
	count := 0
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == '#' {
				count++
			}
		}
	}

	return count
}
