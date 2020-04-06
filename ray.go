package main

// Ray contains a ray data
type Ray struct {
	origin, direction Vec3D
}

// PointAt computes ray position at a time
func (ray Ray) PointAt(time float64) Vec3D {
	return AddVec3D(ray.origin, ScalarProduct(time, ray.direction))
}
