package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"thunder-game/structs"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	title              string = "Thunder"
	windowWidth        int    = 1024
	windowHeight       int    = 1024
	logicalWidth       int    = 800
	logicalHeight      int    = 800
	tileSize           int    = 8
	fireSpreadInterval int    = 3
)

var (
	colorBackground = color.Gray{Y: 128}
	colorFire       = color.RGBA{231, 36, 6, 255}

	boardWidth  = 64
	boardHeight = 64

	boardPixelWidth  = boardWidth * tileSize
	boardPixelHeight = boardHeight * tileSize

	middleOffsetX = (logicalWidth - boardPixelWidth) / 2
	middleOffsetY = (logicalHeight - boardPixelHeight) / 2

	forestDensity = 0.6
)

type Game struct {
	currentLevel *structs.Level
	rng          *rand.Rand
	loop         []GameLoop{}
}

type GameLoop interface {
	Step()
}

func (g *Game) Update() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()

		boardX, boardY := (x-middleOffsetX)/tileSize, (y-middleOffsetY)/tileSize

		if boardX > 0 && boardX <= boardWidth && boardY > 0 && boardY <= boardHeight {
			g.currentLevel.Board[boardY][boardX].Ignite(g.rng)
		}
		fmt.Printf("Mouse clicked at: %d, %d\n", boardX, boardY)

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.currentLevel = structs.InitLevel(boardHeight, boardWidth, forestDensity, g.rng, fireSpreadInterval)
	}

	for _, instance := range g.loop {
		instance.Step()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Thunder InDev!")

	for y, row := range g.currentLevel.Board {
		for x, tile := range row {
			clr := tile.Clr
			if tile.OnFire {
				clr = colorFire
			}

			for i := 0; i < tileSize; i++ {
				for j := 0; j < tileSize; j++ {
					screen.Set(x*tileSize+i+middleOffsetX, y*tileSize+j+middleOffsetY, clr)
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

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	initLevel := structs.InitLevel(boardHeight, boardWidth, forestDensity, rng, fireSpreadInterval)

	if err := ebiten.RunGame(&Game{currentLevel: initLevel, rng: rng, loop: []interface{}{}}); err != nil {
		log.Fatal(err)
	}
}
