package main

type Ray struct {
	origin, direction Vec3D
}

func (ray Ray) PointAt(time float64) Vec3D {
	return AddVec3D(ray.origin, ScalarProduct(time, ray.direction))
}
