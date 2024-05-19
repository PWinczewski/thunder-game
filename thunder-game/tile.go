package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	Barren TileType = iota
	Forest
	Mountain
	Meadow
	Settlement
	Burned
)

type Tile struct {
	TileType           TileType
	Flammable          bool
	OnFire             bool
	Sprite             *ebiten.Image
	IgnitionResistance float64
	DestructionValue   int
	X                  int
	Y                  int
}

func (t *Tile) Ignite(rng *rand.Rand) {
	if t.Flammable && !t.OnFire && rng.Float64() > t.IgnitionResistance {
		t.OnFire = true
		t.Sprite = Sprites["tileBurning"]
	}
}

func initTileForest(x, y int) *Tile {
	return &Tile{Forest, true, false, Sprites["tileForestv2"], 0, 5, x, y}
}

func initTileBarren(x, y int) *Tile {
	return &Tile{Barren, false, false, Sprites["tileBarrenv2"], 1, 1, x, y}
}

// func initTileMountain(x, y int) *Tile {
// 	return &Tile{Mountain, false, false, ColorMountain, 1, 25, x, y}
// }

// func initTileMeadow(x, y int) *Tile {
// 	return &Tile{Meadow, true, false, ColorMeadow, 0, 2, x, y}
// }

// func initTileSettlement(x, y int) *Tile {
// 	return &Tile{Settlement, true, false, ColorSettlement, 0.8, 100, x, y}
// }

func InitTileBurned(x, y int) *Tile {
	return &Tile{Burned, false, false, Sprites["tileBurned"], 1, 1, x, y}
}
