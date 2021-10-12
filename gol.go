package gogol

import (
	"math/rand"
	"time"
)

type Board [][]bool

func getRandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func GenerateBoard(width, height int, randomize bool) Board {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)

		for j := range board[i] {
			if randomize == true {
				board[i][j] = getRandomBool()
			} else {
				board[i][j] = false
			}
		}
	}
	return board
}

func CountLivingNeighborCells(board Board, x, y int) int {
	count := 0

	xRange := []int{x - 1, x, x + 1}
	yRange := []int{y - 1, y, y + 1}

	for _, cy := range yRange {
		if cy < 0 || cy == len(board) {
			// Out of bounds
			continue
		}

		for _, cx := range xRange {
			if (cy == y && cx == x) || cx < 0 || cx == len(board[cy]) {
				// Case 1: current cell is the source cell
				// Case 2: Out of bounds
				continue
			}

			if board[cy][cx] == true {
				count++
				continue
			}
		}
	}

	return count
}

func CalculateEvolution(board Board) Board {
	nextGenBoard := GenerateBoard(len(board[0]), len(board), false)

	for y := range board {
		for x := range board[y] {
			livingNeighbors := CountLivingNeighborCells(board, x, y)

			if livingNeighbors <= 1 {
				// Death by isolation
				nextGenBoard[y][x] = false
				continue
			}

			if livingNeighbors == 2 {
				// Survival
				nextGenBoard[y][x] = board[y][x] == true
				continue
			}

			if livingNeighbors == 3 {
				// Birth
				nextGenBoard[y][x] = true
				continue
			}

			if livingNeighbors >= 4 {
				// Death by overcrowding
				nextGenBoard[y][x] = false
				continue
			}
		}
	}

	return nextGenBoard
}

type Stats struct {
	Alive int `json:"alive"`
	Dead  int `json:"dead"`
	Total int `json:"total"`
}

func CollectStatistics(board Board) Stats {
	stats := Stats{
		Alive: 0,
		Dead:  0,
		Total: 0,
	}

	for y := range board {
		for x := range board[y] {
			stats.Total++

			if board[y][x] == true {
				stats.Alive++
			} else {
				stats.Dead++
			}
		}
	}

	return stats
}
