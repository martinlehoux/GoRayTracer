package main

import "math"

// Camera handle view position and screen
type Camera struct {
	lowerLeftCorner              Vec3D
	origin, horizontal, vertical Vec3D
}

// NewCamera creates a new camera
func NewCamera(vfov float64, aspect float64) Camera {
	origin := Vec3D{0.0, 0.0, 0.0}
	theta := DegreesToRadians(vfov)
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	lowerLeftCorner := Vec3D{-halfWidth, -halfHeight, -1.0}
	horizontal := Vec3D{2.0 * halfWidth, 0.0, 0.0}
	vertical := Vec3D{0.0, 2.0 * halfHeight, 0.0}
	return Camera{lowerLeftCorner, origin, horizontal, vertical}
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
