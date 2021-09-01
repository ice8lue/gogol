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

func countLivingNeighborCells(board Board, x, y int) int {
	count := 0

	xRange := []int{x - 1, x, x + 1}
	yRange := []int{y - 1, y, y + 1}

	for i := range yRange {
		if i < 0 || i > len(board) {
			// Out of bounds
			continue
		}

		for j := range xRange {
			if (i == y && j == x) || j < 0 || j > len(board[i]) {
				// Case 1: current cell is the source cell
				// Case 2: Out of bounds
				continue
			}

			if board[i][j] == true {
				count++
				continue
			}
		}
	}

	return count
}

func CalculateEvolution(board Board) Board {
	for y := range board {
		for x := range board[y] {
			livingNeighbors := countLivingNeighborCells(board, x, y)

			if livingNeighbors <= 1 {
				// Death by isolation
				board[y][x] = false
				continue
			}

			if livingNeighbors == 2 {
				// Survival
				board[y][x] = true
				continue
			}

			if livingNeighbors == 3 {
				// Birth
				board[y][x] = true
				continue
			}

			if livingNeighbors >= 4 {
				// Death by overcrowding
				board[y][x] = false
				continue
			}
		}
	}

	return board
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
