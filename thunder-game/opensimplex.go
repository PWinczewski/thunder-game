package main

import "github.com/ojrac/opensimplex-go"

func getOpensimplexNoise(seed, width, height int64, frequency float64) [][]float64 {
	noise := opensimplex.New(seed)
	heightmap := make([][]float64, height)
	for y := int64(0); y < height; y++ {
		heightmap[y] = make([]float64, width)
		for x := int64(0); x < width; x++ {
			xFloat := float64(x) / float64(width)
			yFloat := float64(y) / float64(height)
			heightmap[y][x] = noise.Eval2(xFloat*frequency, yFloat*frequency)
		}
	}
	return heightmap
}
