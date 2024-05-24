package main

import (
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

type ByLength [][]*Tile

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

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
	l.Board = l.generateBoard(rng)
	l.DestructionTarget = l.getTargetDestruction()
	return l
}

func (l *Level) generateBoard(rng *rand.Rand) [][]*Tile {
	board := make([][]*Tile, l.BoardHeight)
	forestTileCount := int(float64(l.BoardWidth*l.BoardHeight) * l.ForestDensity)
	treeLocations := rng.Perm(l.BoardWidth * l.BoardHeight)[:forestTileCount]

	for i := range board {
		board[i] = make([]*Tile, l.BoardWidth)
		for j := range board[i] {
			board[i][j] = initTileBarren(j, i)
		}
	}

	for _, location := range treeLocations {
		row := location / l.BoardWidth
		col := location % l.BoardWidth
		board[row][col] = initTileForest(col, row)
	}

	return board
}

func (l *Level) getTargetDestruction() int {
	clusters := l.findClusters(Forest)

	clusterSizes := make([]int, 0, len(clusters))
	for _, arr := range clusters {
		size := len(arr)
		if size > 2 {
			clusterSizes = append(clusterSizes, size)
		}
	}

	possibleSums := getPossibleSums(clusterSizes, 2, 3)

	return getBestTarget(possibleSums, 0.1)
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

func combinations(arr []int, X int) [][]int {
	var result [][]int
	var helper func(start int, path []int)

	helper = func(start int, path []int) {
		if len(path) == X {
			comb := make([]int, X)
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(arr); i++ {
			helper(i+1, append(path, arr[i]))
		}
	}

	helper(0, []int{})
	return result
}

func getPossibleSums(arr []int, Nmin int, Nmax int) map[int]int {
	sumMap := make(map[int]int)

	for k := Nmin; k <= Nmax; k++ {
		combs := combinations(arr, k)
		for _, comb := range combs {
			sum := 0
			for _, num := range comb {
				sum += num
			}
			sumMap[sum]++
		}
	}

	return sumMap
}

func getBestTarget(targets map[int]int, tolerance float64) int {
	target := 0
	targetStrength := 0

	for t := range targets {
		strength := 0
		for key, val := range targets {
			if float64(key) <= float64(t)*(1+tolerance) && float64(key) >= float64(t)*(1-tolerance) {
				strength += val
			}
		}
		if strength > targetStrength {
			targetStrength = strength
			target = t
		}
	}
	return target
}
