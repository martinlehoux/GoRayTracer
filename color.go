package main

type Color struct {
	r, g, b float64
}

func (color Color) ToPixel() Pixel {
	return Pixel{int(color.r * 255), int(color.g * 255), int(color.b * 255)}
}
