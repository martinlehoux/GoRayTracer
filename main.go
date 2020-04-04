package main

const HEIGHT = 100
const WIDTH = 200

func main() {
	frame := Frame{}
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			color := Color{
				float64(y) / float64(WIDTH),
				float64(x) / float64(HEIGHT),
				0.2,
			}
			frame[HEIGHT-1-x][y] = color.ToPixel()
		}
	}
	frame.Save("img.ppm")
}
