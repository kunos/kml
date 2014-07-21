package kml

import (
	"fmt"
	"math"
)

type Vec3f struct {
	X float32
	Y float32
	Z float32
}

func (v *Vec3f) Equal(v2 Vec3f) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

func (v *Vec3f) Normalize() {
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))

	if l != 0 {
		v.X /= l
		v.Y /= l
		v.Z /= l
	}
}

func NormalizeVec3f(result, v *Vec3f) *Vec3f {
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))

	if l != 0 {
		result.X = v.X / l
		result.Y = v.Y / l
		result.Z = v.Z / l

	}
	return result
}

func (v Vec3f) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3f) Distance(v2 Vec3f) float32 {

	var t Vec3f

	t.X = v.X - v2.X
	t.Y = v.Y - v2.Y
	t.Z = v.Z - v2.Z

	return float32(math.Sqrt(float64(t.X*t.X + t.Y*t.Y + t.Z*t.Z)))
}

func AddVec3f(result, v1, v2 *Vec3f) *Vec3f {

	result.X = v1.X + v2.X
	result.Y = v1.Y + v2.Y
	result.Z = v1.Z + v2.Z

	return result
}

func SubVec3f(result, v1, v2 *Vec3f) *Vec3f {

	result.X = v1.X - v2.X
	result.Y = v1.Y - v2.Y
	result.Z = v1.Z - v2.Z

	return result
}

func MulVec3f(result, v1, v2 *Vec3f) *Vec3f {

	result.X = v1.X * v2.X
	result.Y = v1.Y * v2.Y
	result.Z = v1.Z * v2.Z

	return result
}

func ScaleVec3f(result, v1 *Vec3f, v float32) *Vec3f {
	result.X = v1.X * v
	result.Y = v1.Y * v
	result.Z = v1.Z * v

	return result
}

func CrossVec3f(result, v1, v2 *Vec3f) *Vec3f {
	result.X = v1.Y*v2.Z - v1.Z*v2.Y
	result.Y = v1.Z*v2.X - v1.X*v2.Z
	result.Z = v1.X*v2.Y - v1.Y*v2.X

	return result
}

func CrossVec3fSlow(v1, v2 *Vec3f) Vec3f {
	var result Vec3f
	result.X = v1.Y*v2.Z - v1.Z*v2.Y
	result.Y = v1.Z*v2.X - v1.X*v2.Z
	result.Z = v1.X*v2.Y - v1.Y*v2.X

	return result

}

func DotVec3f(v1 *Vec3f, v2 *Vec3f) float32 {
	return (((v1.X * v2.X) + (v1.Y * v2.Y)) + (v1.Z * v2.Z))
}

func DivVec3f(result, v1 *Vec3f, f float32) *Vec3f {
	result.X = v1.X / f
	result.Y = v1.Y / f
	result.Z = v1.Z / f

	return result
}

func (v Vec3f) String() string {
	return fmt.Sprintf("[%f,%f,%f]", v.X, v.Y, v.Z)
}

func (v *Vec3f) Add(v1 Vec3f) Vec3f {
	return Vec3f{
		v.X + v1.X,
		v.Y + v1.Y,
		v.Z + v1.Z}
}

func (v *Vec3f) Sub(v1 Vec3f) Vec3f {
	return Vec3f{
		v.X - v1.X,
		v.Y - v1.Y,
		v.Z - v1.Z}
}

// Add In Place
func (v *Vec3f) AddIP(v1 Vec3f) {
	v.X += v1.X
	v.Y += v1.Y
	v.Z += v1.Z
}

func (v *Vec3f) ScaleIP(m float32) {
	v.X *= m
	v.Y *= m
	v.Z *= m
}

func (v *Vec3f) Scale(f float32) Vec3f {
	return Vec3f{
		v.X * f,
		v.Y * f,
		v.Z * f}
}

func (v *Vec3f) ToVec3d() Vec3d {
	return Vec3d{float64(v.X), float64(v.Y), float64(v.Z)}
}

func LerpVec3f(v0, v1 Vec3f, t float32) Vec3f {
	it := 1.0 - t
	return Vec3f{v0.X*it + v1.X*t,
		v0.Y*it + v1.Y*t,
		v0.Z*it + v1.Z*t}
}
