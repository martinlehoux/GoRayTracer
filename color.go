package main

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

func RayColor(ray Ray) Color {
	param := 0.5 * (ray.direction.Unit().y + 1)
	return LinearBlend(Color{0.5, 0.7, 1.0}, Color{1.0, 1.0, 1.0}, param)
}
