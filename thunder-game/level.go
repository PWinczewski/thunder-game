package main

import (
	"math/rand"
)

type Level struct {
	BoardWidth         int
	BoardHeight        int
	ForestDensity      float64
	Board              [][]*Tile
	FireSpreadClock    int
	FireSpreadInterval int
	rng                *rand.Rand
}

func (l *Level) generateBoard(rng *rand.Rand) {
	board := make([][]*Tile, l.BoardHeight)
	for i := range board {
		board[i] = make([]*Tile, l.BoardWidth)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = initTileBarren(j, i)
		}
	}

	nums := rng.Perm(l.BoardWidth * l.BoardHeight)[:int(float64(l.BoardWidth*l.BoardHeight)*l.ForestDensity)]

	for _, num := range nums {
		row := num / l.BoardHeight
		col := num % l.BoardWidth
		board[row][col] = initTileForest(col, row)
	}

	l.Board = board
}

func (l *Level) Step() {
	if l.FireSpreadInterval == l.FireSpreadClock {
		l.FireSpreadClock = 0
		for _, t := range l.GetTilesOnFire() {
			if t.OnFire {
				if t.Y-1 >= 0 {
					l.Board[t.Y-1][t.X].Ignite(l.rng)
				}
				if t.X+1 < len(l.Board[0]) {
					l.Board[t.Y][t.X+1].Ignite(l.rng)
				}
				if t.Y+1 < len(l.Board) {
					l.Board[t.Y+1][t.X].Ignite(l.rng)
				}
				if t.X-1 >= 0 {
					l.Board[t.Y][t.X-1].Ignite(l.rng)
				}
				l.Board[t.Y][t.X] = InitTileBurned(t.X, t.Y)
			}
		}
	}

	l.FireSpreadClock++
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

func InitLevel(boardWidth int, boardHeight int, forestDensity float64, rng *rand.Rand, fireSpreadInterval int) *Level {
	l := &Level{
		BoardWidth:         boardWidth,
		BoardHeight:        boardHeight,
		ForestDensity:      forestDensity,
		FireSpreadInterval: fireSpreadInterval,
		rng:                rng,
	}
	l.generateBoard(rng)
	return l
}
