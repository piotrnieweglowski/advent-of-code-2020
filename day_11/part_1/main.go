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
			if board[r][c] == 'L' && adjacentSeatsTaken(board, r, c) == 0 {
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
			if board[r][c] == '#' && adjacentSeatsTaken(board, r, c) >= 4 {
				row = append(row, 'L')
			} else {
				row = append(row, board[r][c])
			}
		}
		newBoard = append(newBoard, row)
	}

	return newBoard
}

func adjacentSeatsTaken(board [][]byte, row int, column int) int {
	return isSeatTaken(board, row-1, column-1) +
		isSeatTaken(board, row-1, column) +
		isSeatTaken(board, row-1, column+1) +
		isSeatTaken(board, row, column-1) +
		isSeatTaken(board, row, column+1) +
		isSeatTaken(board, row+1, column-1) +
		isSeatTaken(board, row+1, column) +
		isSeatTaken(board, row+1, column+1)
}

func isSeatTaken(board [][]byte, row int, column int) int {
	rows := len(board)
	columns := len(board[0])

	if row < 0 || row >= rows {
		return 0
	}

	if column < 0 || column >= columns {
		return 0
	}

	if board[row][column] == '#' {
		return 1
	}

	return 0
}

func anyPlaceToTake(board [][]byte) bool {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 'L' && adjacentSeatsTaken(board, r, c) == 0 {
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
