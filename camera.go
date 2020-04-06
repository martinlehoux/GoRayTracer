package main

// Camera handle view position and screen
type Camera struct {
	lowerLeftCorner              Vec3D
	origin, horizontal, vertical Vec3D
}

// GetRay computes Ray given screen params
func (camera Camera) GetRay(u float64, v float64) Ray {
	return Ray{
		camera.origin,
		AddVec3D(AddVec3D(
			camera.lowerLeftCorner,
			ScalarProduct(u, camera.horizontal)),
			ScalarProduct(v, camera.vertical),
		),
	}
}
