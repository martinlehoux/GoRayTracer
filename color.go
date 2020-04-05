package main

import "math"

type Color struct {
	r, g, b float64
}

func (color Color) ToPixel() Pixel {
	return Pixel{int(color.r * 255), int(color.g * 255), int(color.b * 255)}
}

func LinearBlend(color1 Color, color2 Color, param float64) Color {
	return Color{
		param*color1.r + (1.0-param)*color2.r,
		param*color1.g + (1.0-param)*color2.g,
		param*color1.b + (1.0-param)*color2.b,
	}
}

func AttenuateColor(attenuation Color, color Color) Color {
	return Color{
		attenuation.r * color.r,
		attenuation.g * color.g,
		attenuation.b * color.b,
	}
}

func RayColor(ray Ray, hitable_list HitableList, depth int) Color {
	if depth > MAX_DEPTH {
		return Color{0.0, 0.0, 0.0}
	}
	hit := hitable_list.Hit(ray, T_MIN, T_MAX)
	if hit.time > 0.0 {
		scattered, attenuation := hit.material.Scatter(ray, hit)
		return AttenuateColor(attenuation, RayColor(scattered, hitable_list, depth+1))
	}
	param := 0.5 * (ray.direction.Unit().y + 1)
	return LinearBlend(Color{0.5, 0.7, 1.0}, Color{1.0, 1.0, 1.0}, param)
}

func ColorMean(color_list [RAY_PER_PIXEL]Color) Color {
	length := float64(len(color_list))
	mean_color := Color{}
	for _, color := range color_list {
		mean_color.r += color.r
		mean_color.g += color.g
		mean_color.b += color.b
	}
	mean_color.r /= length
	mean_color.g /= length
	mean_color.b /= length
	return mean_color
}

func ColorMeanSquare(color_list [RAY_PER_PIXEL]Color) Color {
	mean_color := ColorMean(color_list)
	mean_color.r = math.Sqrt(mean_color.r)
	mean_color.g = math.Sqrt(mean_color.g)
	mean_color.b = math.Sqrt(mean_color.b)
	return mean_color
}
