package main

import (
	"math"
	"math/rand"
)

type Vec3D struct {
	x, y, z float64
}

func (vec Vec3D) Invert() Vec3D {
	return Vec3D{-vec.x, -vec.y, -vec.z}
}

func (vec Vec3D) Length() float64 {
	return math.Sqrt(DotProduct(vec, vec))
}

func (vec Vec3D) Unit() Vec3D {
	length := vec.Length()
	if length == 0 {
		return vec
	}
	return ScalarProduct(1.0/length, vec)
}

func AddVec3D(vec1 Vec3D, vec2 Vec3D) Vec3D {
	return Vec3D{vec1.x + vec2.x, vec1.y + vec2.y, vec1.z + vec2.z}
}

func SubVec3D(vec1 Vec3D, vec2 Vec3D) Vec3D {
	return AddVec3D(vec1, vec2.Invert())
}

func ScalarProduct(scalar float64, vec Vec3D) Vec3D {
	return Vec3D{scalar * vec.x, scalar * vec.y, scalar * vec.z}
}

func DotProduct(vec1 Vec3D, vec2 Vec3D) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y + vec1.z*vec2.z
}

func CrossProduct(vec1 Vec3D, vec2 Vec3D) Vec3D {
	return Vec3D{
		vec1.y*vec2.z - vec1.z*vec2.y,
		vec1.z*vec2.x - vec1.x*vec2.z,
		vec1.x*vec2.y - vec1.y*vec2.x,
	}
}

func RandomInUnitSphere() Vec3D {
	for {
		point := Vec3D{2*rand.Float64() - 1, 2*rand.Float64() - 1, 2*rand.Float64() - 1}
		if point.Length() < 1.0 {
			return point
		}
	}
}
