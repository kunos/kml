package kml

import "math"

type Vec2f struct {
	X float32
	Y float32
}

func (v Vec2f) Equal(v2 Vec2f) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v *Vec2f) Normalize() float32 {
	l := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))

	if l != 0 {
		v.X /= l
		v.Y /= l
	}

	return l
}

func (v *Vec2f) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v1 *Vec2f) Add(v2 Vec2f) Vec2f {
	return Vec2f{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 *Vec2f) Sub(v2 Vec2f) Vec2f {
	return Vec2f{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vec2f) Scale(f float32) Vec2f {
	return Vec2f{v1.X * f, v1.Y * f}
}

func (v1 *Vec2f) Mul(v2 Vec2f) Vec2f {
	return Vec2f{v1.X * v2.X, v1.Y * v2.Y}
}

func (v1 *Vec2f) Div(f float32) Vec2f {
	return Vec2f{v1.X / f, v1.Y / f}
}

func AddVec2f(result, v1, v2 *Vec2f) *Vec2f {

	result.X = v1.X + v1.X
	result.Y = v1.Y + v1.Y

	return result
}

func SubVec2f(result, v1, v2 *Vec2f) *Vec2f {

	result.X = v1.X - v1.X
	result.Y = v1.Y - v1.Y

	return result
}

func DistanceVec2f(v1, v2 *Vec2f) float32 {
	lx := v1.X - v2.X
	ly := v1.Y - v2.Y

	return float32(math.Sqrt(float64(lx*lx + ly*ly)))

}

func DotVec2f(v1 *Vec2f, v2 *Vec2f) float32 {
	return ((v1.X * v2.X) + (v1.Y * v2.Y))
}
