package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	initLevel := InitLevel(boardWidth, boardHeight, forestDensity, rng, fireSpreadInterval, destructionTargetTolerance)

	return &Game{currentLevel: initLevel, rng: rng, loop: []GameLoop{initLevel}}
}

func clickedOnBoard(x, y int) (int, int, bool) {

	boardX, boardY := (x-BoardOffsetX)/tileSize, (y-BoardOffsetY)/tileSize

	onBoard := boardX >= 0 && boardX < boardWidth && boardY >= 0 && boardY < boardHeight

	return boardX, boardY, onBoard
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		x, y := ebiten.CursorPosition()

		if bx, by, ok := clickedOnBoard(x, y); ok {
			g.currentLevel.Board[by][bx].Ignite()
			fmt.Printf("Mouse clicked at: %d, %d\n", bx, by)
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		Fullscreen = !Fullscreen
		ebiten.SetFullscreen(Fullscreen)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.currentLevel = InitLevel(boardWidth, boardHeight, forestDensity, g.rng, fireSpreadInterval, destructionTargetTolerance)
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
			op.GeoM.Translate(float64(BoardOffsetX+x*tileSize), float64(BoardOffsetY+y*tileSize))
			screen.DrawImage(sprite, op)
		}
	}

	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.Filter = ebiten.FilterLinear

	TextGUIDestruction := strconv.Itoa(g.currentLevel.Destruction) + "/" + strconv.Itoa(g.currentLevel.DestructionTarget)

	text.Draw(screen, TextGUIDestruction, &text.GoTextFace{
		Source: Fonts["Alkhemikal"],
		Size:   FontSizes["medium"],
	}, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return internalWidth, internalHeight
}
