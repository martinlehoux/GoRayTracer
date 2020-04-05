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

type Metal struct {
	albedo Color
	fuzz   float64
}

func (metal Metal) Scatter(ray Ray, hit HitRecord) (Ray, Color) {
	reflected := AddVec3D(ray.direction.Unit().Reflect(hit.normal), ScalarProduct(metal.fuzz, RandomInUnitSphere()))
	scattered := Ray{hit.point, reflected}
	return scattered, metal.albedo
}
