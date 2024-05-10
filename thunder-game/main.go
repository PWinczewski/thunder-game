package main

import (
	"fmt"
	"image/color"
	"log"
	"thunder-game/structs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	title         string = "Thunder"
	windowWidth   int    = 1024
	windowHeight  int    = 1024
	logicalWidth  int    = 800
	logicalHeight int    = 800
	tileSize      int    = 16
)

var (
	colorBackground = color.Gray{Y: 128}
	colorFire       = color.RGBA{231, 36, 6, 255}
)

type Game struct {
	currentLevel       *structs.Level
	fireSpreadInterval int
	fireSpreadClock    int
}

func (g *Game) Update() error {
	g.fireSpreadClock++

	boardPixelWidth := len(g.currentLevel.Board[0]) * tileSize
	boardPixelHeight := len(g.currentLevel.Board) * tileSize
	offsetX := (logicalWidth - boardPixelWidth) / 2
	offsetY := (logicalHeight - boardPixelHeight) / 2

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()

		boardX, boardY := (x-offsetX)/tileSize, (y-offsetY)/tileSize

		g.currentLevel.Board[boardY][boardX].Ignite()
		fmt.Printf("Mouse clicked at: %d, %d\n", boardX, boardY)
	}
	if g.fireSpreadInterval == g.fireSpreadClock {
		g.fireSpreadClock = 0
		for _, t := range g.currentLevel.GetTilesOnFire() {
			if t.OnFire {
				if t.Y-1 >= 0 {
					g.currentLevel.Board[t.Y-1][t.X].Ignite()
				}
				if t.X+1 < len(g.currentLevel.Board[0]) {
					g.currentLevel.Board[t.Y][t.X+1].Ignite()
				}
				if t.Y+1 < len(g.currentLevel.Board) {
					g.currentLevel.Board[t.Y+1][t.X].Ignite()
				}
				if t.X-1 >= 0 {
					g.currentLevel.Board[t.Y][t.X-1].Ignite()
				}
				g.currentLevel.Board[t.Y][t.X] = structs.InitTileBurned(t.X, t.Y)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Thunder InDev!")

	//Centering jargon
	boardWidth := len(g.currentLevel.Board[0]) * tileSize
	boardHeight := len(g.currentLevel.Board) * tileSize
	offsetX := (logicalWidth - boardWidth) / 2
	offsetY := (logicalHeight - boardHeight) / 2

	for y, row := range g.currentLevel.Board {
		for x, tile := range row {
			var clr color.Color

			if tile.OnFire {
				clr = colorFire
			} else {
				clr = tile.Clr
			}

			for i := 0; i < tileSize; i++ {
				for j := 0; j < tileSize; j++ {
					screen.Set(x*tileSize+i+offsetX, y*tileSize+j+offsetY, clr)
				}
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return logicalWidth, logicalHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Thunder InDev")

	initLevel := structs.NewLevel(32, 32, 0.5)
	fireSpreadInterval := 15

	if err := ebiten.RunGame(&Game{currentLevel: initLevel, fireSpreadInterval: fireSpreadInterval}); err != nil {
		log.Fatal(err)
	}
}
