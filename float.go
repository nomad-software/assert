package assert

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Approx64 returns true if the passed 64bit floats are within the passed
// epsilon value of each other.
func approx64(a float64, b float64, epsilon float64) bool {
	if a == b {
		return true
	}

	minNormal := math.Float64frombits(0x0010000000000000)
	maxFloat := float64(math.MaxFloat64)

	absA := abs(a)
	absB := abs(b)
	diff := abs(a - b)

	if a == 0 || b == 0 || absA+absB < minNormal {
		return diff < epsilon*minNormal
	}

	return diff/min(absA+absB, maxFloat) < epsilon
}

// Approx32 returns true if the passed 32bit floats are within the passed
// epsilon value of each other.
func approx32(a float32, b float32, epsilon float32) bool {
	if a == b {
		return true
	}

	minNormal := math.Float32frombits(0x00800000)
	maxFloat := float32(math.MaxFloat32)

	absA := abs(a)
	absB := abs(b)
	diff := abs(a - b)

	if a == 0 || b == 0 || absA+absB < minNormal {
		return diff < epsilon*minNormal
	}

	return diff/min(absA+absB, maxFloat) < epsilon
}

// Abs returns the absolute value of the passed float.
func abs[T constraints.Float](val T) T {
	switch {
	case val < 0:
		return -val
	case val == 0:
		return 0
	}
	return val
}

// Min returns the minimum of the two passed values.
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
