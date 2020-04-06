package main

import "math"

// HitRecord contains data of a ray hiting a hitable
type HitRecord struct {
	time     float64
	point    Vec3D
	normal   Vec3D
	material Material
}

// Hitable is a physical volume a Ray can hit
type Hitable interface {
	Hit(ray Ray, tmin float64, tmax float64) HitRecord
	Normal(point Vec3D, ray Ray) Vec3D
}

// HitableList ...
type HitableList []Hitable

// Hit finds the closest hit of all hitables
func (hitableList HitableList) Hit(ray Ray, tmin float64, tmax float64) HitRecord {
	closestHit := HitRecord{TMax, Vec3D{}, Vec3D{}, nil}
	for _, hitable := range hitableList {
		hit := hitable.Hit(ray, tmin, closestHit.time)
		if hit.time > 0 {
			closestHit = hit
		}
	}
	if closestHit.time < TMax {
		return closestHit
	}
	return HitRecord{}
}

// Sphere ...
type Sphere struct {
	center   Vec3D
	radius   float64
	material Material
}

// Normal ...
func (sphere Sphere) Normal(point Vec3D, ray Ray) Vec3D {
	return SubVec3D(point, sphere.center).Unit()
}

// Hit ...
func (sphere Sphere) Hit(ray Ray, tmin float64, tmax float64) HitRecord {
	oc := SubVec3D(ray.origin, sphere.center)
	a := DotProduct(ray.direction, ray.direction)
	b := 2.0 * DotProduct(oc, ray.direction)
	c := DotProduct(oc, oc) - sphere.radius*sphere.radius
	discriminant := b*b - 4*a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		time := (-b - root) / (2.0 * a)
		if tmin < time && time < tmax {
			point := ray.PointAt(time)
			return HitRecord{time, point, sphere.Normal(point, ray), sphere.material}
		}
		time = (-b + root) / (2.0 * a)
		if tmin < time && time < tmax {
			point := ray.PointAt(time)
			return HitRecord{time, point, sphere.Normal(point, ray), sphere.material}
		}
	}
	return HitRecord{}
}
