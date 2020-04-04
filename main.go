package main

const HEIGHT = 100
const WIDTH = 200

func main() {
	LOWER_LEFT_CORNER := Vec3D{-2.0, -1.0, -1.0}
	ORIGIN := Vec3D{0.0, 0.0, 0.0}
	HORIZONTAL := Vec3D{4.0, 0.0, 0.0}
	VERTICAL := Vec3D{0.0, 2.0, 0.0}
	SPHERE := Sphere{Vec3D{0.0, 0.0, -1.0}, .5}
	frame := Frame{}
	for x := 0; x < HEIGHT; x++ {
		for y := 0; y < WIDTH; y++ {
			u, v := float64(y)/float64(WIDTH), float64(x)/float64(HEIGHT)
			ray := Ray{ORIGIN, AddVec3D(AddVec3D(LOWER_LEFT_CORNER, ScalarProduct(u, HORIZONTAL)), ScalarProduct(v, VERTICAL))}
			color := RayColor(ray, SPHERE)
			frame[HEIGHT-1-x][y] = color.ToPixel()
		}
	}
	frame.Save("img.ppm")
}
