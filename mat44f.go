package kml

import (
	"errors"
	"fmt"
	"math"
)

type Mat44f struct {
	M11 float32
	M12 float32
	M13 float32
	M14 float32

	M21 float32
	M22 float32
	M23 float32
	M24 float32

	M31 float32
	M32 float32
	M33 float32
	M34 float32

	M41 float32
	M42 float32
	M43 float32
	M44 float32
}

func (matrix *Mat44f) Ortho(left float32, top float32, right float32, bottom float32, zNearPlane float32, zFarPlane float32) {

	matrix.M11 = 2.0 / (right - left)
	//matrix.M12 = matrix.M13 = matrix.M14 = 0.0f
	matrix.M22 = 2.0 / (top - bottom)
	//matrix.M21 = matrix.M23 = matrix.M24 = 0.0f;
	matrix.M33 = 1.0 / (zNearPlane - zFarPlane)
	//matrix.M31 = matrix.M32 = matrix.M34 = 0.0
	matrix.M41 = (left + right) / (left - right)
	matrix.M42 = (top + bottom) / (bottom - top)
	matrix.M43 = zNearPlane / (zNearPlane - zFarPlane)
	matrix.M44 = 1.0

}

func CreatePerspective(fieldOfView, aspectRatio, nearPlaneDistance, farPlaneDistance float32) Mat44f {
	res := Mat44f{}

	res.Perspective(fieldOfView, aspectRatio, nearPlaneDistance, farPlaneDistance)

	return res
}

func CreateIdentityMatrix() Mat44f {

	return Mat44f{1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0}

}

func CreateTarget(pos, target Vec3f) Mat44f {

	forward := target.Sub(pos)
	forward.Normalize()

	var right Vec3f

	upv := Vec3f{0, 1, 0}

	CrossVec3f(&right, &forward, &upv)

	var up Vec3f
	CrossVec3f(&up, &right, &forward)

	var res Mat44f
	right.Normalize()
	up.Normalize()
	res.SetForward(&forward)
	res.SetRight(&right)
	res.SetUP(&up)

	res.SetTranslation(pos)

	return res

}

func (m *Mat44f) Perspective(fieldOfView, aspectRatio, nearPlaneDistance, farPlaneDistance float32) {

	num := float32(1.0 / (math.Tan(float64(fieldOfView) * 0.5)))
	num9 := num / aspectRatio

	m.M11 = num9
	m.M12 = 0
	m.M13 = 0
	m.M14 = 0.0
	m.M22 = num
	m.M21 = 0
	m.M23 = 0
	m.M24 = 0.0
	m.M31 = 0
	m.M32 = 0.0
	//m.M33 = farPlaneDistance / (nearPlaneDistance - farPlaneDistance) // XNA
	m.M33 = -(farPlaneDistance + nearPlaneDistance) / (farPlaneDistance - nearPlaneDistance) // GLM
	m.M34 = -1.0

	m.M41 = 0
	m.M42 = 0
	//m.M43 = (nearPlaneDistance * farPlaneDistance) / (nearPlaneDistance - farPlaneDistance) // XNA
	m.M43 = -(2.0 * nearPlaneDistance * farPlaneDistance) / (farPlaneDistance - nearPlaneDistance) // GLM
	m.M44 = 0

}

func (m *Mat44f) SetIdentity() {
	m.M11 = 1.0
	m.M12 = 0
	m.M13 = 0
	m.M14 = 0
	m.M21 = 0.0
	m.M22 = 1.0
	m.M23 = 0
	m.M24 = 0
	m.M31 = 0.0
	m.M32 = 0
	m.M33 = 1.0
	m.M34 = 0
	m.M41 = 0.0
	m.M42 = 0
	m.M43 = 0
	m.M44 = 1.0
}

func CreateScaleMat44f(scale Vec3f) Mat44f {
	var matrix Mat44f

	matrix.M11 = scale.X
	matrix.M22 = scale.Y
	matrix.M33 = scale.Z
	matrix.M44 = 1.0

	return matrix
}

func CreateOrtho(left, right, top, bottom, znear, zfar float32) Mat44f {
	matrix := Mat44f{}
	matrix.M11 = 2.0 / (right - left)
	matrix.M22 = 2.0 / (top - bottom)
	matrix.M33 = 1.0 / (znear - zfar)
	matrix.M41 = (left + right) / (left - right)
	matrix.M42 = (top + bottom) / (bottom - top)
	matrix.M43 = znear / (znear - zfar)
	matrix.M44 = 1.0
	return matrix
}

func CreateRotationMatrix4(axis Vec3f, angle float32) Mat44f {
	angle = TO_RAD * angle

	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	t := 1 - c

	var res Mat44f
	res.M11 = c + axis.X*axis.X*t
	res.M22 = c + axis.Y*axis.Y*t
	res.M33 = c + axis.Z*axis.Z*t
	res.M44 = 1.0

	tmp1 := axis.X * axis.Y * t
	tmp2 := axis.Z * s
	res.M12 = tmp1 + tmp2
	res.M21 = tmp1 - tmp2

	tmp1 = axis.X * axis.Z * t
	tmp2 = axis.Y * s
	res.M13 = tmp1 - tmp2
	res.M31 = tmp1 + tmp2

	tmp1 = axis.Y * axis.Z * t
	tmp2 = axis.X * s
	res.M23 = tmp1 + tmp2
	res.M32 = tmp1 - tmp2

	return res
	/*return Mat44f{axis.X*axis.X*k + c, axis.X*axis.Y*k + axis.Z*s, axis.X*axis.Z*k - axis.Y*s, 0,
	axis.X*axis.Y*k - axis.Z*s, axis.Y*axis.Y*k + c, axis.Y*axis.Z*k + axis.X*s, 0,
	axis.X*axis.Z*k + axis.Y*s, axis.Y*axis.Z*k - axis.X*s, axis.Z*axis.Z*k + c, 0,
	0, 0, 0, 1}*/
}

/*
func CreateLookAtMatrix(pos, dir, up *Vec3f) Mat44f {
	matrix:=Mat44f{}
	var vector2 Vec3f
	CrossVec3f(&vector2, up, dir)
	vector2.Normalize()

	vector2.normalize();
		vec3f vector3 = cross(vector, vector2);
		matrix.M11 = vector2.x;
		matrix.M12 = vector3.x;
		matrix.M13 = vector.x;
		matrix.M14 = 0.0f;
		matrix.M21 = vector2.y;
		matrix.M22 = vector3.y;
		matrix.M23 = vector.y;
		matrix.M24 = 0.0f;
		matrix.M31 = vector2.z;
		matrix.M32 = vector3.z;
		matrix.M33 = vector.z;
		matrix.M34 = 0.0f;
		matrix.M41 = -dot(vector2, cameraPosition);
		matrix.M42 = -dot(vector3, cameraPosition);
		matrix.M43 = -dot(vector, cameraPosition);
		matrix.M44 = 1.0f;
		return matrix;
}*/

func (m *Mat44f) View() Mat44f {
	res := Mat44f{}

	pos := Vec3f{m.M41, m.M42, m.M43}

	up := Vec3f{m.M21, m.M22, m.M23}
	forw := Vec3f{m.M31, m.M32, m.M33}
	up.Normalize()
	forw.Normalize()
	var right Vec3f
	CrossVec3f(&right, &up, &forw)
	right.Normalize()

	res.M11 = right.X
	res.M12 = up.X
	res.M13 = forw.X

	res.M21 = right.Y
	res.M22 = up.Y
	res.M23 = forw.Y

	res.M31 = right.Z
	res.M32 = up.Z
	res.M33 = forw.Z

	res.M41 = -DotVec3f(&right, &pos)
	res.M42 = -DotVec3f(&up, &pos)
	res.M43 = -DotVec3f(&forw, &pos)

	res.M44 = 1.0
	return res
}
func (m *Mat44f) SetTranslation(t Vec3f) {
	m.M41 = t.X
	m.M42 = t.Y
	m.M43 = t.Z
}

func (m *Mat44f) Translation() Vec3f {
	return Vec3f{m.M41, m.M42, m.M43}
}

func (m *Mat44f) SetForward(value *Vec3f) {
	m.M31 = -value.X
	m.M32 = -value.Y
	m.M33 = -value.Z
}

func (m *Mat44f) Forward() Vec3f {
	return Vec3f{-m.M31, -m.M32, -m.M33}
}

func (m *Mat44f) Right() Vec3f {
	return Vec3f{m.M11, m.M12, m.M13}
}

func (m *Mat44f) SetRight(value *Vec3f) {
	m.M11 = value.X
	m.M12 = value.Y
	m.M13 = value.Z
}

func (m *Mat44f) Up() Vec3f {
	return Vec3f{m.M21, m.M22, m.M23}
}

func (m *Mat44f) SetUP(value *Vec3f) {
	m.M21 = value.X
	m.M22 = value.Y
	m.M23 = value.Z
}

func (m1 *Mat44f) Mul(m2 *Mat44f) Mat44f {
	return Mat44f{
		m1.M11*m2.M11 + m1.M21*m2.M12 + m1.M31*m2.M13 + m1.M41*m2.M14,
		m1.M12*m2.M11 + m1.M22*m2.M12 + m1.M32*m2.M13 + m1.M42*m2.M14,
		m1.M13*m2.M11 + m1.M23*m2.M12 + m1.M33*m2.M13 + m1.M43*m2.M14,
		m1.M14*m2.M11 + m1.M24*m2.M12 + m1.M34*m2.M13 + m1.M44*m2.M14,

		m1.M11*m2.M21 + m1.M21*m2.M22 + m1.M31*m2.M23 + m1.M41*m2.M24,
		m1.M12*m2.M21 + m1.M22*m2.M22 + m1.M32*m2.M23 + m1.M42*m2.M24,
		m1.M13*m2.M21 + m1.M23*m2.M22 + m1.M33*m2.M23 + m1.M43*m2.M24,
		m1.M14*m2.M21 + m1.M24*m2.M22 + m1.M34*m2.M23 + m1.M44*m2.M24,

		m1.M11*m2.M31 + m1.M21*m2.M32 + m1.M31*m2.M33 + m1.M41*m2.M34,
		m1.M12*m2.M31 + m1.M22*m2.M32 + m1.M32*m2.M33 + m1.M42*m2.M34,
		m1.M13*m2.M31 + m1.M23*m2.M32 + m1.M33*m2.M33 + m1.M43*m2.M34,
		m1.M14*m2.M31 + m1.M24*m2.M32 + m1.M34*m2.M33 + m1.M44*m2.M34,

		m1.M11*m2.M41 + m1.M21*m2.M42 + m1.M31*m2.M43 + m1.M41*m2.M44,
		m1.M12*m2.M41 + m1.M22*m2.M42 + m1.M32*m2.M43 + m1.M42*m2.M44,
		m1.M13*m2.M41 + m1.M23*m2.M42 + m1.M33*m2.M43 + m1.M43*m2.M44,
		m1.M14*m2.M41 + m1.M24*m2.M42 + m1.M34*m2.M43 + m1.M44*m2.M44}

}

func (m *Mat44f) MulVec4(vec Vec3f) Vec3f {
	tmp := Vec3f{}
	tmp.X = vec.X*m.M11 + vec.Y*m.M21 + vec.Z*m.M31 + m.M41
	tmp.Y = vec.X*m.M12 + vec.Y*m.M22 + vec.Z*m.M32 + m.M42
	tmp.Z = vec.X*m.M13 + vec.Y*m.M23 + vec.Z*m.M33 + m.M43

	return tmp
}

func (m *Mat44f) Invert() (Mat44f, error) {
	tmp := Mat44f{}
	det := m.Determinant()
	if det == 0 {
		return tmp, errors.New("non-invertible matrix")
	}

	tmp.M11 = m.M32*m.M43*m.M24 - m.M42*m.M33*m.M24 + m.M42*m.M23*m.M34 - m.M22*m.M43*m.M34 - m.M32*m.M23*m.M44 + m.M22*m.M33*m.M44
	tmp.M21 = m.M41*m.M33*m.M24 - m.M31*m.M43*m.M24 - m.M41*m.M23*m.M34 + m.M21*m.M43*m.M34 + m.M31*m.M23*m.M44 - m.M21*m.M33*m.M44
	tmp.M31 = m.M31*m.M42*m.M24 - m.M41*m.M32*m.M24 + m.M41*m.M22*m.M34 - m.M21*m.M42*m.M34 - m.M31*m.M22*m.M44 + m.M21*m.M32*m.M44
	tmp.M41 = m.M41*m.M32*m.M23 - m.M31*m.M42*m.M23 - m.M41*m.M22*m.M33 + m.M21*m.M42*m.M33 + m.M31*m.M22*m.M43 - m.M21*m.M32*m.M43
	tmp.M12 = m.M42*m.M33*m.M14 - m.M32*m.M43*m.M14 - m.M42*m.M13*m.M34 + m.M12*m.M43*m.M34 + m.M32*m.M13*m.M44 - m.M12*m.M33*m.M44
	tmp.M22 = m.M31*m.M43*m.M14 - m.M41*m.M33*m.M14 + m.M41*m.M13*m.M34 - m.M11*m.M43*m.M34 - m.M31*m.M13*m.M44 + m.M11*m.M33*m.M44
	tmp.M32 = m.M41*m.M32*m.M14 - m.M31*m.M42*m.M14 - m.M41*m.M12*m.M34 + m.M11*m.M42*m.M34 + m.M31*m.M12*m.M44 - m.M11*m.M32*m.M44
	tmp.M42 = m.M31*m.M42*m.M13 - m.M41*m.M32*m.M13 + m.M41*m.M12*m.M33 - m.M11*m.M42*m.M33 - m.M31*m.M12*m.M43 + m.M11*m.M32*m.M43
	tmp.M13 = m.M22*m.M43*m.M14 - m.M42*m.M23*m.M14 + m.M42*m.M13*m.M24 - m.M12*m.M43*m.M24 - m.M22*m.M13*m.M44 + m.M12*m.M23*m.M44
	tmp.M23 = m.M41*m.M23*m.M14 - m.M21*m.M43*m.M14 - m.M41*m.M13*m.M24 + m.M11*m.M43*m.M24 + m.M21*m.M13*m.M44 - m.M11*m.M23*m.M44
	tmp.M33 = m.M21*m.M42*m.M14 - m.M41*m.M22*m.M14 + m.M41*m.M12*m.M24 - m.M11*m.M42*m.M24 - m.M21*m.M12*m.M44 + m.M11*m.M22*m.M44
	tmp.M43 = m.M41*m.M22*m.M13 - m.M21*m.M42*m.M13 - m.M41*m.M12*m.M23 + m.M11*m.M42*m.M23 + m.M21*m.M12*m.M43 - m.M11*m.M22*m.M43
	tmp.M14 = m.M32*m.M23*m.M14 - m.M22*m.M33*m.M14 - m.M32*m.M13*m.M24 + m.M12*m.M33*m.M24 + m.M22*m.M13*m.M34 - m.M12*m.M23*m.M34
	tmp.M24 = m.M21*m.M33*m.M14 - m.M31*m.M23*m.M14 + m.M31*m.M13*m.M24 - m.M11*m.M33*m.M24 - m.M21*m.M13*m.M34 + m.M11*m.M23*m.M34
	tmp.M34 = m.M31*m.M22*m.M14 - m.M21*m.M32*m.M14 - m.M31*m.M12*m.M24 + m.M11*m.M32*m.M24 + m.M21*m.M12*m.M34 - m.M11*m.M22*m.M34
	tmp.M44 = m.M21*m.M32*m.M13 - m.M31*m.M22*m.M13 + m.M31*m.M12*m.M23 - m.M11*m.M32*m.M23 - m.M21*m.M12*m.M33 + m.M11*m.M22*m.M33

	inv_det := 1.0 / det
	tmp.M11 = tmp.M11 * inv_det
	tmp.M21 = tmp.M21 * inv_det
	tmp.M31 = tmp.M31 * inv_det
	tmp.M41 = tmp.M41 * inv_det
	tmp.M12 = tmp.M12 * inv_det
	tmp.M22 = tmp.M22 * inv_det
	tmp.M32 = tmp.M32 * inv_det
	tmp.M42 = tmp.M42 * inv_det
	tmp.M13 = tmp.M13 * inv_det
	tmp.M23 = tmp.M23 * inv_det
	tmp.M33 = tmp.M33 * inv_det
	tmp.M43 = tmp.M43 * inv_det
	tmp.M14 = tmp.M14 * inv_det
	tmp.M24 = tmp.M24 * inv_det
	tmp.M34 = tmp.M34 * inv_det
	tmp.M44 = tmp.M44 * inv_det

	return tmp, nil
}

// The determinant of this matrix.
func (m *Mat44f) Determinant() float32 {
	return m.M14*m.M23*m.M32*m.M41 -
		m.M13*m.M24*m.M32*m.M41 -
		m.M14*m.M22*m.M33*m.M41 +
		m.M12*m.M24*m.M33*m.M41 +
		m.M13*m.M22*m.M34*m.M41 -
		m.M12*m.M23*m.M34*m.M41 -
		m.M14*m.M23*m.M31*m.M41 +
		m.M13*m.M24*m.M31*m.M41 +
		m.M14*m.M21*m.M33*m.M41 -
		m.M11*m.M24*m.M33*m.M41 -
		m.M13*m.M21*m.M34*m.M41 +
		m.M11*m.M23*m.M34*m.M41 +
		m.M14*m.M22*m.M31*m.M43 -
		m.M12*m.M24*m.M31*m.M43 -
		m.M14*m.M21*m.M32*m.M43 +
		m.M11*m.M24*m.M32*m.M43 +
		m.M12*m.M21*m.M34*m.M43 -
		m.M11*m.M22*m.M34*m.M43 -
		m.M13*m.M22*m.M31*m.M44 +
		m.M12*m.M23*m.M31*m.M44 +
		m.M13*m.M21*m.M32*m.M44 -
		m.M11*m.M23*m.M32*m.M44 -
		m.M12*m.M21*m.M33*m.M44 +
		m.M11*m.M22*m.M33*m.M44
}

func (m Mat44f) String() string {
	return fmt.Sprintf("%f %f %f %f\n%f %f %f %f\n%f %f %f %f\n%f %f %f %f\n",
		m.M11, m.M21, m.M31, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		m.M14, m.M24, m.M34, m.M44)

}

func (matrix *Mat44f) TransformVec3f(v Vec3f, homo bool) Vec3f {
	var res Vec3f

	num3 := (((v.X * matrix.M11) + (v.Y * matrix.M21)) + (v.Z * matrix.M31)) + matrix.M41
	num2 := (((v.X * matrix.M12) + (v.Y * matrix.M22)) + (v.Z * matrix.M32)) + matrix.M42
	num := (((v.X * matrix.M13) + (v.Y * matrix.M23)) + (v.Z * matrix.M33)) + matrix.M43

	if homo {

		w := (((v.X * matrix.M14) + (v.Y * matrix.M24)) + (v.Z * matrix.M34)) + matrix.M44

		res.X = num3 / w
		res.Y = num2 / w
		res.Z = num / w

	} else {
		res.X = num3
		res.Y = num2
		res.Z = num
	}

	return res
}
