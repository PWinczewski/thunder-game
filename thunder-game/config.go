package main

import (
	"image/color"
)

const (
	title              string = "Thunder"
	windowWidth        int    = 1024
	windowHeight       int    = 1024
	logicalWidth       int    = 800
	logicalHeight      int    = 800
	tileSize           int    = 4
	fireSpreadInterval int    = 3
)

var (
	colorBackground = color.Gray{Y: 128}
	colorFire       = color.RGBA{231, 36, 6, 255}

	boardWidth  = 20
	boardHeight = 50

	boardPixelWidth  = boardWidth * tileSize
	boardPixelHeight = boardHeight * tileSize

	middleOffsetX = (logicalWidth - boardPixelWidth) / 2
	middleOffsetY = (logicalHeight - boardPixelHeight) / 2

	forestDensity = 0.6
)
