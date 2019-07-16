package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

/*
The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in
one of two possible states, alive or dead, (or populated and unpopulated, respectively). Every cell interacts with its
eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent.

At each step in time, the following transitions occur:
- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with three live neighbours becomes a live cell, as if by reproduction.

The initial pattern constitutes the seed of the system. The first generation is created by applying the above rules
simultaneously to every cell in the seed; births and deaths occur simultaneously, and the discrete moment at which this
happens is sometimes called a tick. Each generation is a pure function of the preceding one.

The rules continue to be applied repeatedly to create further generations.
*/

type state int

const (
	dead  state = 0
	alive state = 1

	displayAlive string = ` `
	displayDead  string = `█`
)

var (
	// Oscillator - Beacon
	seed = [][]state{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
)

func main() {
	clear()

	genLength := len(seed)
	rowLength := len(seed[0])

	// seed the board
	board := seed

	// malloc the next generation
	next := make([][]state, genLength)
	for i := range next {
		next[i] = make([]state, rowLength)
	}

	/*
		Implementation edges rules:
			y == 0 => up is dead
			y == len(board) - 1 => down is dead
			x == 0 => left is dead
			x == len(row) - 1 => right is dead
	*/
	for {
		for y, row := range board {
			for x, cell := range row {
				neighbours := 0

				// count alive neighbours
				switch y {
				case 0: // up is dead
					switch x {
					case 0: // left is dead
						neighbours = countAliveNeighbours(board[y][x+1], board[y+1][x], board[y+1][x+1])
					case rowLength - 1: // right is dead
						neighbours = countAliveNeighbours(board[y][x-1], board[y+1][x], board[y+1][x-1])
					default:
						neighbours = countAliveNeighbours(board[y][x-1], board[y+1][x-1], board[y+1][x], board[y+1][x+1], board[y][x+1])
					}
				case genLength - 1: // down is dead
					switch x {
					case 0: // left is dead
						neighbours = countAliveNeighbours(board[y-1][x], board[y-1][x+1], board[y][x+1])
					case rowLength - 1: // right is dead
						neighbours = countAliveNeighbours(board[y-1][x], board[y-1][x-1], board[y][x-1])
					default:
						neighbours = countAliveNeighbours(board[y][x-1], board[y-1][x-1], board[y-1][x], board[y-1][x+1], board[y][x+1])
					}
				default:
					switch x {
					case 0: // left is dead
						neighbours = countAliveNeighbours(board[y-1][x], board[y-1][x+1], board[y][x+1], board[y+1][x+1], board[y+1][x])
					case rowLength - 1: // right is dead
						neighbours = countAliveNeighbours(board[y-1][x], board[y-1][x-1], board[y][x-1], board[y+1][x-1], board[y+1][x])
					default:
						neighbours = countAliveNeighbours(board[y-1][x-1], board[y-1][x], board[y-1][x+1], board[y][x+1], board[y+1][x+1], board[y+1][x], board[y+1][x-1], board[y][x-1])
					}
				}

				// rules of life
				var future state
				switch cell {
				case alive:
					switch neighbours {
					case 2, 3:
						future = alive
					default:
						future = dead
					}

					printAlive()
				case dead:
					switch neighbours {
					case 3:
						future = alive
					default:
						future = dead
					}

					printDead()
				}
				next[y][x] = future

				if x == rowLength-1 {
					printBreak()
				}
			}
		}

		// regenerate board with next generation
		regenerate(next, board)

		tick()
		clear()
	}
}

func printBreak() {
	fmt.Println()
}

func printDead() {
	fmt.Print(displayDead)
}

func printAlive() {
	fmt.Print(displayAlive)
}

func regenerate(next [][]state, board [][]state) {
	for y, row := range next {
		for x, cell := range row {
			board[y][x] = cell
		}
	}
}

func tick() {
	time.Sleep(1 * time.Second)
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	_ = c.Run()
}

func countAliveNeighbours(n ...state) int {
	c := 0

	for _, v := range n {
		if v == alive {
			c++
		}
	}

	return c
}
