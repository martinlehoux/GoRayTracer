package main

import "math"

type HitRecord struct {
	time   float64
	point  Vec3D
	normal Vec3D
}

type Hitable interface {
	Hit(ray Ray, t_min float64, t_max float64) HitRecord
	Normal(point Vec3D, ray Ray) Vec3D
}

type HitableList []Hitable

func (hitable_list HitableList) Hit(ray Ray, t_min float64, t_max float64) HitRecord {
	closest_hit := HitRecord{T_MAX, Vec3D{}, Vec3D{}}
	for _, hitable := range hitable_list {
		hit := hitable.Hit(ray, t_min, closest_hit.time)
		if hit.time > 0 {
			closest_hit = hit
		}
	}
	if closest_hit.time < T_MAX {
		return closest_hit
	} else {
		return HitRecord{}
	}
}

type Sphere struct {
	center Vec3D
	radius float64
}

func (sphere Sphere) Normal(point Vec3D, ray Ray) Vec3D {
	out_normal := SubVec3D(point, sphere.center).Unit()
	if DotProduct(ray.direction, out_normal) < 0.0 {
		return out_normal
	} else {
		return out_normal.Invert()
	}
}

func (sphere Sphere) Hit(ray Ray, t_min float64, t_max float64) HitRecord {
	oc := SubVec3D(ray.origin, sphere.center)
	a := DotProduct(ray.direction, ray.direction)
	b := 2.0 * DotProduct(oc, ray.direction)
	c := DotProduct(oc, oc) - sphere.radius*sphere.radius
	discriminant := b*b - 4*a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		time := (-b - root) / (2.0 * a)
		if t_min < time && time < t_max {
			point := ray.PointAt(time)
			return HitRecord{time, point, sphere.Normal(point, ray)}
		}
		time = (-b + root) / (2.0 * a)
		if t_min < time && time < t_max {
			point := ray.PointAt(time)
			return HitRecord{time, point, sphere.Normal(point, ray)}
		}
	}
	return HitRecord{}
}
