package main

const HEIGHT = 100
const WIDTH = 200

func main() {
	frame := Frame{}
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			frame[HEIGHT-1-x][y] = Pixel{
				r: 255 * y / WIDTH,
				g: 255 * x / HEIGHT,
				b: 255 * 0.2,
			}
		}
	}
	frame.Save("img.ppm")
}
