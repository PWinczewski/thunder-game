package main

import (
	"math"
	"math/rand"
)

// Point represents a point in 2D space
type Point struct {
	X, Y float64
}

// EuclideanDistance calculates the Euclidean distance between two points
func EuclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// GenerateFeaturePoints generates random feature points within the given dimensions
func GenerateFeaturePoints(numPoints, width, height int) []Point {
	points := make([]Point, numPoints)
	for i := range points {
		points[i] = Point{
			X: rand.Float64() * float64(width),
			Y: rand.Float64() * float64(height),
		}
	}
	return points
}

// WorleyNoise generates a Worley noise texture and returns it as a 2D array
func getWorleyNoise(width, height, numPoints int) [][]float64 {
	featurePoints := GenerateFeaturePoints(numPoints, width, height)
	noise := make([][]float64, height)
	for i := range noise {
		noise[i] = make([]float64, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			minDist := math.MaxFloat64
			for _, point := range featurePoints {
				dist := EuclideanDistance(Point{float64(x), float64(y)}, point)
				if dist < minDist {
					minDist = dist
				}
			}
			noise[y][x] = minDist
		}
	}

	// Normalize the noise values to the range [0, 1]
	maxDist := 0.0
	for y := range noise {
		for x := range noise[y] {
			if noise[y][x] > maxDist {
				maxDist = noise[y][x]
			}
		}
	}

	for y := range noise {
		for x := range noise[y] {
			noise[y][x] /= maxDist
		}
	}

	return noise
}
