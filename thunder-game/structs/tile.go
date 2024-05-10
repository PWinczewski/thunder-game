package structs

import "image/color"

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
	ColorBurned     = color.RGBA{33, 4, 67, 255}
)

type Tile struct {
	TileType         TileType
	Flammable        bool
	OnFire           bool
	Clr              color.RGBA
	DestructionValue int
	X                int
	Y                int
}

func (t *Tile) Ignite() {
	if t.Flammable && !t.OnFire {
		t.OnFire = true
	}
}

func initTileForest(x, y int) *Tile {
	return &Tile{Forest, true, false, ColorForest, 5, x, y}
}

func initTileBarren(x, y int) *Tile {
	return &Tile{Barren, false, false, ColorBarren, 1, x, y}
}

func initTileMountain(x, y int) *Tile {
	return &Tile{Mountain, false, false, ColorMountain, 25, x, y}
}

func initTileMeadow(x, y int) *Tile {
	return &Tile{Meadow, true, false, ColorMeadow, 2, x, y}
}

func initTileSettlement(x, y int) *Tile {
	return &Tile{Settlement, true, false, ColorSettlement, 100, x, y}
}

func InitTileBurned(x, y int) *Tile {
	return &Tile{Burned, false, false, ColorBurned, 1, x, y}
}
