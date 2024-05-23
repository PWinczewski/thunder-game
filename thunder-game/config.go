package main

import (
	"image/color"
)

const (
	title              string = "Thunder"
	windowWidth        int    = 1280
	windowHeight       int    = 720
	internalWidth      int    = 640
	internalHeight     int    = 360
	tileSize           int    = 16
	fireSpreadInterval int    = 8
)

var (
	colorBackground = color.Gray{Y: 128}

	boardWidth  = 32
	boardHeight = 16

	boardPixelWidth  = boardWidth * tileSize
	boardPixelHeight = boardHeight * tileSize

	middleBoardOffsetX = (internalWidth - boardPixelWidth) / 2
	middleBoardOffsetY = (internalHeight - boardPixelHeight) / 2

	forestDensity    = 0.5
	SpreadDirections = []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	Fullscreen = false
)
