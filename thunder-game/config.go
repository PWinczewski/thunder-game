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

	boardWidth  = 40
	boardHeight = 18

	boardPixelWidth  = boardWidth * tileSize
	boardPixelHeight = boardHeight * tileSize

	BoardOffsetX = (internalWidth - boardPixelWidth) / 2
	BoardOffsetY = (internalHeight-boardPixelHeight)/2 - (internalHeight-tileSize*boardHeight)/2

	forestDensity    = 0.5
	SpreadDirections = []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	destructionTargetTolerance = 0.15

	maxStrikes = 3

	FontSizes = map[string]float64{"small": 16, "medium": 32, "large": 48}

	Fullscreen = false
)
