package main

type Material interface {
	Scatter(ray Ray, hit HitRecord) (Ray, Color)
}

type Lambertian struct {
	albedo Color
}

func (lambertian Lambertian) Scatter(ray Ray, hit HitRecord) (Ray, Color) {
	// TODO: Use the correct distribution
	direction := AddVec3D(hit.normal, RandomInUnitSphere())
	scattered := Ray{hit.point, direction}
	return scattered, lambertian.albedo
}
