package main

import (
	"math"
	"math/rand"
)

// Material generate color when hitable is hit
type Material interface {
	Scatter(ray Ray, hit HitRecord) (Ray, Color)
}

// Lambertian is diffuse material
type Lambertian struct {
	albedo Color
}

// Scatter for Lambertians
func (lambertian Lambertian) Scatter(ray Ray, hit HitRecord) (Ray, Color) {
	// TODO: Use the correct distribution
	var normal Vec3D
	if DotProduct(ray.direction, hit.normal) < 0.0 {
		normal = hit.normal
	} else {
		normal = hit.normal.Invert()
	}
	direction := AddVec3D(normal, RandomInUnitSphere())
	scattered := Ray{hit.point, direction}
	return scattered, lambertian.albedo
}

// Metal material
type Metal struct {
	albedo Color
	fuzz   float64
}

// Scatter for Metal
func (metal Metal) Scatter(ray Ray, hit HitRecord) (Ray, Color) {
	var normal Vec3D
	if DotProduct(ray.direction, hit.normal) < 0.0 {
		normal = hit.normal
	} else {
		normal = hit.normal.Invert()
	}
	reflected := AddVec3D(ray.direction.Unit().Reflect(normal), ScalarProduct(metal.fuzz, RandomInUnitSphere()))
	scattered := Ray{hit.point, reflected}
	return scattered, metal.albedo
}

// Dielectric is glass for instance
type Dielectric struct {
	eta float64
}

// Scatter for Dielectric
func (dielectric Dielectric) Scatter(ray Ray, hit HitRecord) (Ray, Color) {
	var etaI, etaT float64
	var normal Vec3D
	attenuation := Color{1.0, 1.0, 1.0}
	if DotProduct(ray.direction, hit.normal) < 0.0 {
		etaI, etaT = 1.0, dielectric.eta
		normal = hit.normal
	} else {
		etaI, etaT = 1.0, dielectric.eta
		normal = hit.normal.Invert()
	}
	vecI := ray.direction.Unit()
	cosTheta := FMin(DotProduct(vecI.Invert(), normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaI/etaT*sinTheta > 1.0 || rand.Float64() < Schlick(cosTheta, etaI, etaT) {
		reflected := ray.direction.Unit().Reflect(normal)
		scattered := Ray{hit.point, reflected}
		return scattered, attenuation
	}
	refracted := Refract(vecI, normal, etaI, etaT)
	scattered := Ray{hit.point, refracted}
	return scattered, attenuation
}

// Schlick glass refraction probability
func Schlick(cosine float64, etaI float64, etaT float64) float64 {
	r0 := (1.0 - etaI/etaT) / (1.0 + etaI/etaT)
	r0 *= r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5)
}
