// TODO Use pointers when needed to reduce memory
package main

import (
	"math/rand"

	"github.com/schollz/progressbar"
)

const (
	// Height of frame
	Height = 1000
	// Width of frame
	Width = 2000
	// TMin is the minimal no zero time
	TMin = 0.001
	// TMax limits time
	TMax = 1000
	// RaysPerPixel for anti-aliasing
	RaysPerPixel = 10
	// MaxDepth for ray boucing
	MaxDepth = 50
	// FPS Frames per second for movies
	FPS = 60
)

func main() {
	// Scene
	lowerLeftCorner := Vec3D{-2.0, -1.0, -1.0}
	origin := Vec3D{0.0, 0.0, 0.0}
	horizontal := Vec3D{4.0, 0.0, 0.0}
	vertical := Vec3D{0.0, 2.0, 0.0}
	camera := Camera{lowerLeftCorner, origin, horizontal, vertical}

	matRed := Lambertian{Color{0.7, 0.3, 0.3}}
	matGreen := Lambertian{Color{0.8, 0.8, 0.0}}
	metalGrey := Metal{Color{0.8, 0.8, 0.8}, 0.3}
	// metalBrown := Metal{Color{0.8, 0.6, 0.2}, 1.0}
	glass := Dielectric{1.5}

	WORLD := HitableList{
		Sphere{Vec3D{0.0, 0.0, -1.0}, 0.5, glass},
		Sphere{Vec3D{-1.0, 0.0, -1.0}, 0.5, metalGrey},
		Sphere{Vec3D{1.0, 0.0, -1.0}, 0.5, matRed},
		Sphere{Vec3D{0.0, -100.5, -1.0}, 100.0, matGreen},
	}

	// Program
	// start := AddVec3D(Origin, Vec3D{-1.0, 0.0, 0.0})
	// move := Vec3D{2.0, 0.0, 0.0}
	frame := Frame{}
	bar := progressbar.New(Height)
	bar.RenderBlank()
	// for k := 0; k < FPS*DURATION; k++ {
	// origin := AddVec3D(start, ScalarProduct(float64(k)/float64(FPS*DURATION), move))
	// camera.origin = origin
	for x := 0; x < Height; x++ {
		for y := 0; y < Width; y++ {
			var colorList [RaysPerPixel]Color
			for it := 0; it < RaysPerPixel; it++ {
				u := (float64(y) + rand.Float64()) / float64(Width)
				v := (float64(x) + rand.Float64()) / float64(Height)
				ray := camera.GetRay(u, v)
				colorList[it] = RayColor(ray, WORLD, 0)
			}
			color := ColorMeanSquare(colorList)
			frame[Height-1-x][y] = color.ToPixel()
		}
		bar.Add(1)
	}
	frame.Save("img.png")
	// frame.Save(fmt.Sprintf("frames/frame_%06d.ppm", k))
}

// }
