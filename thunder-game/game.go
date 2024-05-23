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
	count        int
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

	initLevel := InitLevel(boardWidth, boardHeight, forestDensity, rng, fireSpreadInterval)

	return &Game{currentLevel: initLevel, rng: rng, loop: []GameLoop{initLevel}}
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()

		boardX, boardY := (x-middleBoardOffsetX)/tileSize, (y-middleBoardOffsetY)/tileSize

		if boardX >= 0 && boardX < boardWidth && boardY >= 0 && boardY < boardHeight {
			g.currentLevel.Board[boardY][boardX].Ignite()
		}
		fmt.Printf("Mouse clicked at: %d, %d\n", boardX, boardY)

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		Fullscreen = !Fullscreen
		ebiten.SetFullscreen(Fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.currentLevel = InitLevel(boardWidth, boardHeight, forestDensity, g.rng, fireSpreadInterval)
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
			sprite := tile.Sprite

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(middleBoardOffsetX+x*tileSize), float64(middleBoardOffsetY+y*tileSize))
			screen.DrawImage(sprite, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return internalWidth, internalHeight
}
