package main

import "math"

// Color in RGB, each val being between 0 and 1
type Color struct {
	r, g, b float64
}

// ToPixel converts a 0-1 color to a 0-255 pixel
func (color Color) ToPixel() Pixel {
	return Pixel{int(color.r * 255), int(color.g * 255), int(color.b * 255)}
}

// LinearBlend mixes two colors with a float param
func LinearBlend(color1 Color, color2 Color, param float64) Color {
	return Color{
		param*color1.r + (1.0-param)*color2.r,
		param*color1.g + (1.0-param)*color2.g,
		param*color1.b + (1.0-param)*color2.b,
	}
}

// AttenuateColor multiplies two colors
func AttenuateColor(attenuation Color, color Color) Color {
	return Color{
		attenuation.r * color.r,
		attenuation.g * color.g,
		attenuation.b * color.b,
	}
}

// RayColor finds a hit for a ray, scatters the ray and return the ray color
func RayColor(ray Ray, hitableList HitableList, depth int) Color {
	if depth > MaxDepth {
		return Color{0.0, 0.0, 0.0}
	}
	hit := hitableList.Hit(ray, TMin, TMax)
	if hit.time > 0.0 {
		scattered, attenuation := hit.material.Scatter(ray, hit)
		return AttenuateColor(attenuation, RayColor(scattered, hitableList, depth+1))
	}
	param := 0.5 * (ray.direction.Unit().y + 1)
	return LinearBlend(Color{0.5, 0.7, 1.0}, Color{1.0, 1.0, 1.0}, param)
}

// ColorMean computes the mean color of a color list
func ColorMean(colorList [RaysPerPixel]Color) Color {
	length := float64(len(colorList))
	meanColor := Color{}
	for _, color := range colorList {
		meanColor.r += color.r
		meanColor.g += color.g
		meanColor.b += color.b
	}
	meanColor.r /= length
	meanColor.g /= length
	meanColor.b /= length
	return meanColor
}

// ColorMeanSquare is a more realistic color mean
func ColorMeanSquare(colorList [RaysPerPixel]Color) Color {
	meanColor := ColorMean(colorList)
	meanColor.r = math.Sqrt(meanColor.r)
	meanColor.g = math.Sqrt(meanColor.g)
	meanColor.b = math.Sqrt(meanColor.b)
	return meanColor
}
