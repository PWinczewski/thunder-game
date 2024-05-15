package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Thunder InDev")

	if err := ebiten.RunGame(InitGame()); err != nil {
		log.Fatal(err)
	}
}
