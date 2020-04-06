package main

import "math"

// FMin return mins of floats
func FMin(a float64, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * 2.0 * math.Pi / 360.0
}
