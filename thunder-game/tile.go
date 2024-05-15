package main

import (
	"image/color"
	"math/rand"
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

var (
	ColorForest     = color.RGBA{5, 90, 36, 255}
	ColorMountain   = color.RGBA{65, 60, 75, 255}
	ColorMeadow     = color.RGBA{30, 200, 75, 255}
	ColorBarren     = color.RGBA{67, 58, 46, 255}
	ColorSettlement = color.RGBA{90, 80, 46, 255}
	ColorBurned     = color.RGBA{50, 4, 67, 255}
)

type Tile struct {
	TileType           TileType
	Flammable          bool
	OnFire             bool
	Clr                color.RGBA
	IgnitionResistance float64
	DestructionValue   int
	X                  int
	Y                  int
}

func (t *Tile) Ignite(rng *rand.Rand) {
	if t.Flammable && !t.OnFire && rng.Float64() > t.IgnitionResistance {
		t.OnFire = true
	}
}

func initTileForest(x, y int) *Tile {
	return &Tile{Forest, true, false, ColorForest, 0, 5, x, y}
}

func initTileBarren(x, y int) *Tile {
	return &Tile{Barren, false, false, ColorBarren, 1, 1, x, y}
}

func initTileMountain(x, y int) *Tile {
	return &Tile{Mountain, false, false, ColorMountain, 1, 25, x, y}
}

func initTileMeadow(x, y int) *Tile {
	return &Tile{Meadow, true, false, ColorMeadow, 0, 2, x, y}
}

func initTileSettlement(x, y int) *Tile {
	return &Tile{Settlement, true, false, ColorSettlement, 0.8, 100, x, y}
}

func InitTileBurned(x, y int) *Tile {
	return &Tile{Burned, false, false, ColorBurned, 1, 1, x, y}
}
