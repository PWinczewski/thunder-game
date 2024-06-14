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

	forestDensity = 0.0 // for random approach

	openSimplexThreshold        = 0.0
	destructionTargetTolerance  = 0.1
	minimumClusterSizeForTarget = 5
	noiseFrequency              = 15.0

	ErodePatterns = [][3][3]int64{
		{
			{-1, 0, -1},
			{1, 1, 1},
			{-1, 0, -1},
		},
		{
			{-1, 1, -1},
			{0, 1, 0},
			{-1, 1, -1},
		},
		{
			{-1, 0, -1},
			{0, 1, 0},
			{-1, 0, -1},
		},
	}

	maxStrikes = 5

	FontSizes = map[string]float64{"small": 16, "medium": 32, "large": 48}

	FullscreenToggle = false
)
