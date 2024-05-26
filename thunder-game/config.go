package main

import (
	"image/color"
)

const (
	title              string = "Thunder"
	windowWidth        int    = 1280
	windowHeight       int    = 720
	internalWidth      int    = 1280
	internalHeight     int    = 720
	tileSize           int    = 16
	fireSpreadInterval int    = 8
)

var (
	colorBackground = color.Gray{Y: 128}

	boardWidth  = 80
	boardHeight = 36

	boardPixelWidth  = boardWidth * tileSize
	boardPixelHeight = boardHeight * tileSize

	BoardOffsetX = (internalWidth - boardPixelWidth) / 2
	BoardOffsetY = (internalHeight-boardPixelHeight)/2 - (internalHeight-tileSize*boardHeight)/2

	SpreadDirections = []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	forestDensity               = 0.0
	noiseThreshold              = 0.1
	destructionTargetTolerance  = 0.15
	minimumClusterSizeForTarget = 5
	noiseFrequency              = 10.0

	maxStrikes = 3

	FontSizes = map[string]float64{"small": 16, "medium": 32, "large": 48}

	Fullscreen = false
)
