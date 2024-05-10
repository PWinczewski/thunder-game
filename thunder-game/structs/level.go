package structs

import (
	"math/rand"
	"time"
)

type Level struct {
	boardWidth    int
	boardHeight   int
	forestDensity float64
	Board         [][]*Tile
}

func (l *Level) generateBoard() {
	rand.Seed(time.Now().UnixNano())

	board := make([][]*Tile, l.boardHeight)
	for i := range board {
		board[i] = make([]*Tile, l.boardWidth)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = initTileBarren(j, i)
		}
	}

	// Generate a set of unique random numbers
	nums := rand.Perm(l.boardWidth * l.boardHeight)[:int(float64(l.boardWidth*l.boardHeight)*l.forestDensity)]

	// Assign structs to random indices in the 2D array
	for _, num := range nums {
		row := num / l.boardHeight
		col := num % l.boardWidth
		board[row][col] = initTileForest(col, row)
	}

	l.Board = board
}

func (l *Level) GetTilesOnFire() []*Tile {
	var tilesOnFire []*Tile
	for _, row := range l.Board {
		for _, tile := range row {
			if tile.OnFire {
				tilesOnFire = append(tilesOnFire, tile)
			}
		}
	}
	return tilesOnFire
}

func NewLevel(boardWidth int, boardHeight int, forestDensity float64) *Level {
	l := &Level{
		boardWidth:    boardWidth,
		boardHeight:   boardHeight,
		forestDensity: forestDensity,
	}
	l.generateBoard()
	return l
}
