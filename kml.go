package kml

import (
	"math"
)

const TO_RAD float32 = 0.0174532925
const TO_DEG float32 = 57.2957795

func Sinf(f float32) float32 {
	return float32(math.Sin(float64(f)))
}

func Cosf(f float32) float32 {
	return float32(math.Cos(float64(f)))
}

func Fabs(f float32) float32 {
	return float32(math.Abs(float64(f)))
}

func Sign(f float32) float32 {
	if f > 0 {
		return 1.0
	}

	if f < 0 {
		return -1.0
	}

	return 0.0
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Deadzone(v, dzone float32) float32 {
	if Fabs(v) < dzone {
		return 0
	}

	return v
}

func Clamp(v, min, max float32) float32 {
	if v > max {
		return max
	}

	if v < min {
		return min
	}

	return v
}

func Maxf(v0, v1 float32) float32 {
	if v0 > v1 {
		return v0
	}

	return v1
}

func Minf(v0, v1 float32) float32 {
	if v0 < v1 {
		return v0
	}

	return v1
}

func Saturatef(v float32) float32 {

	if v < 0.0 {
		return 0.0
	}

	if v > 1.0 {
		return 1.0
	}

	return v

}

func Gamma(v, gamma float32) float32 {
	return Fabs(float32(math.Pow(float64(v), float64(gamma)))) * Sign(v)
}
