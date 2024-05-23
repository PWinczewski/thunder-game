package main

import (
	"fmt"
	"math/rand"
)

type Level struct {
	BoardWidth         int
	BoardHeight        int
	ForestDensity      float64
	DestructionTarget  int
	Destruction        int
	Board              [][]*Tile
	FireSpreadClock    int
	FireSpreadInterval int
	rng                *rand.Rand
}

func (l *Level) Step() {
	if l.FireSpreadInterval == l.FireSpreadClock {
		l.FireSpreadClock = 0
		for _, t := range l.GetTilesOnFire() {
			if t.OnFire {
				for _, dir := range SpreadDirections {
					if tile, ok := l.SafeArrayAccess(t.X+dir.dx, t.Y+dir.dy); ok {
						tile.Ignite()
					}
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

func (l *Level) SafeArrayAccess(x, y int) (*Tile, bool) {
	if x >= 0 && x < l.BoardWidth && y >= 0 && y < l.BoardHeight {
		return l.Board[y][x], true
	}
	return nil, false
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

func (l *Level) generateBoard(rng *rand.Rand) {
	board := make([][]*Tile, l.BoardHeight)
	forestTileCount := int(float64(l.BoardWidth*l.BoardHeight) * l.ForestDensity)

	for i := range board {
		board[i] = make([]*Tile, l.BoardWidth)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = initTileBarren(j, i)
		}
	}

	nums := rng.Perm(l.BoardWidth * l.BoardHeight)[:forestTileCount]

	for _, num := range nums {
		row := num / l.BoardWidth
		col := num % l.BoardWidth
		board[row][col] = initTileForest(col, row)
	}

	l.Board = board

	clusters := l.findClusters(Forest)
	fmt.Println(clusters)
}

func (l *Level) findClusters(tileType TileType) [][]*Tile {
	visited := make([][]bool, l.BoardHeight)
	for i := range visited {
		visited[i] = make([]bool, l.BoardWidth)
	}

	var clusters [][]*Tile

	var dfs func(x, y int, cluster *[]*Tile)
	dfs = func(x, y int, cluster *[]*Tile) {
		if x < 0 || y < 0 || x >= l.BoardWidth || y >= l.BoardHeight || visited[y][x] || l.Board[y][x].TileType != tileType {
			return
		}
		visited[y][x] = true
		*cluster = append(*cluster, l.Board[y][x])
		for _, dir := range SpreadDirections {
			dfs(x+dir.dx, y+dir.dy, cluster)
		}

	}

	for i := range l.BoardHeight {
		for j := range l.BoardWidth {
			if !visited[i][j] && l.Board[i][j].TileType == tileType {
				var cluster []*Tile
				dfs(j, i, &cluster)
				if len(cluster) > 0 {
					clusters = append(clusters, cluster)
				}
			}
		}
	}
	return clusters
}
