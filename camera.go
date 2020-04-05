package main

type Camera struct {
	lower_left_corner            Vec3D
	origin, horizontal, vertical Vec3D
}

func (camera Camera) GetRay(u float64, v float64) Ray {
	return Ray{
		camera.origin,
		AddVec3D(AddVec3D(
			camera.lower_left_corner,
			ScalarProduct(u, camera.horizontal)),
			ScalarProduct(v, camera.vertical),
		),
	}
}
