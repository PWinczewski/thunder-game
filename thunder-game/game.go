package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	currentLevel *Level
	rng          *rand.Rand
	loop         []GameLoop
}

type GameLoop interface {
	Step()
}

func InitGame() *Game {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	initLevel := InitLevel(boardHeight, boardWidth, forestDensity, rng, fireSpreadInterval)

	return &Game{currentLevel: initLevel, rng: rng, loop: []GameLoop{initLevel}}
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
		g.currentLevel = InitLevel(boardHeight, boardWidth, forestDensity, g.rng, fireSpreadInterval)
		g.loop = []GameLoop{g.currentLevel}
	}

	for _, instance := range g.loop {
		instance.Step()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Thunder InDev!")

	screen.Fill(colorBackground)
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
