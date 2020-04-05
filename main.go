package main

import (
	"math/rand"

	"github.com/schollz/progressbar"
)

const HEIGHT = 100
const WIDTH = 200
const T_MIN = 0.001
const T_MAX = 1000
const RAY_PER_PIXEL = 100
const MAX_DEPTH = 50

func main() {
	LOWER_LEFT_CORNER := Vec3D{-2.0, -1.0, -1.0}
	ORIGIN := Vec3D{0.0, 0.0, 0.0}
	HORIZONTAL := Vec3D{4.0, 0.0, 0.0}
	VERTICAL := Vec3D{0.0, 2.0, 0.0}
	CAMERA := Camera{LOWER_LEFT_CORNER, ORIGIN, HORIZONTAL, VERTICAL}

	SPHERE := Sphere{Vec3D{0.0, 0.0, -1.0}, 0.5}
	EARTH := Sphere{Vec3D{0.0, -100.5, -1.0}, 100.0}

	WORLD := HitableList{EARTH, SPHERE}
	frame := Frame{}
	bar := progressbar.New(HEIGHT)
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			var color_list [RAY_PER_PIXEL]Color
			for it := 0; it < RAY_PER_PIXEL; it++ {
				u := (float64(y) + rand.Float64()) / float64(WIDTH)
				v := (float64(x) + rand.Float64()) / float64(HEIGHT)
				ray := CAMERA.GetRay(u, v)
				color_list[it] = RayColor(ray, WORLD, 0)
			}
			color := ColorMeanSquare(color_list)
			frame[HEIGHT-1-x][y] = color.ToPixel()
		}
		bar.Add(1)
	}
	frame.Save("img.ppm")
}
