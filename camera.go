package main

import "math"

// Camera handle view position and screen
type Camera struct {
	lowerLeftCorner              Vec3D
	origin, horizontal, vertical Vec3D
}

// NewCamera creates a new camera
func NewCamera(lookFrom Vec3D, lookAt Vec3D, vewUp Vec3D, vfov float64, aspect float64) Camera {
	origin := lookFrom
	theta := DegreesToRadians(vfov)
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	w := SubVec3D(lookFrom, lookAt).Unit()
	u := CrossProduct(vewUp, w)
	v := CrossProduct(w, u)
	lowerLeftCorner := SubVec3D(
		SubVec3D(
			origin,
			AddVec3D(
				ScalarProduct(halfWidth, u),
				ScalarProduct(halfHeight, v),
			),
		),
		w,
	)
	horizontal := ScalarProduct(2.0*halfWidth, u)
	vertical := ScalarProduct(2.0*halfHeight, v)
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
