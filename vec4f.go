package kml

type Vec4f struct {
	X float32
	Y float32
	Z float32
	W float32
}

func AddVec4f(v1 *Vec4f, v2 *Vec4f) (res Vec4f) {

	res.X = v1.X + v2.X
	res.Y = v1.Y + v2.Y
	res.Z = v1.Z + v2.Z
	res.W = v1.W + v2.W

	return

}

func SubVec4f(v1 *Vec4f, v2 *Vec4f) (res Vec4f) {
	res.X = v1.X - v2.X
	res.Y = v1.Y - v2.Y
	res.Z = v1.Z - v2.Z
	res.W = v1.W - v2.W

	return
}

func MulVec4f(v1 *Vec4f, v2 *Vec4f) (res Vec4f) {
	res.X = v1.X * v2.X
	res.Y = v1.Y * v2.Y
	res.Z = v1.Z * v2.Z
	res.W = v1.W * v2.W

	return
}

func ScaleVec4(v1 *Vec4f, f float32) (res Vec4f) {
	res.X = v1.X * f
	res.Y = v1.Y * f
	res.Z = v1.Z * f
	res.W = v1.W * f

	return
}
