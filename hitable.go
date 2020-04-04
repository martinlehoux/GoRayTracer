package main

import "math"

type HitRecord struct {
	time   float64
	point  Vec3D
	normal Vec3D
}

type Hitable interface {
	Hit(ray Ray) HitRecord
	Normal(point Vec3D) Vec3D
}

type Sphere struct {
	center Vec3D
	radius float64
}

func (sphere Sphere) Normal(point Vec3D) Vec3D {
	return (SubVec3D(point, sphere.center)).Unit()
}

func (sphere Sphere) Hit(ray Ray) HitRecord {
	oc := SubVec3D(ray.origin, sphere.center)
	a := DotProduct(ray.direction, ray.direction)
	b := 2.0 * DotProduct(oc, ray.direction)
	c := DotProduct(oc, oc) - sphere.radius*sphere.radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return HitRecord{}
	}
	time := (-b - math.Sqrt(discriminant)) / 2.0 / a
	point := ray.PointAt(time)
	return HitRecord{time, point, SubVec3D(point, sphere.Normal(point))}
}
