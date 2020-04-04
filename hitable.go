package main

type Hitable interface {
	Hit(ray Ray) bool
}

type Sphere struct {
	center Vec3D
	radius float64
}

func (sphere Sphere) Hit(ray Ray) bool {
	oc := SubVec3D(ray.origin, sphere.center)
	a := DotProduct(ray.direction, ray.direction)
	b := 2.0 * DotProduct(oc, ray.direction)
	c := DotProduct(oc, oc) - sphere.radius*sphere.radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
