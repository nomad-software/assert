package assert

import (
	"regexp"
	"testing"

	"golang.org/x/exp/constraints"
)

// True is a helper function to test a boolean value is true.
func True(t *testing.T, val bool) {
	t.Helper()

	if !val {
		t.Errorf("failed asserting true\n")
	}
}

// True is a helper function to test a boolean value is false.
func False(t *testing.T, val bool) {
	t.Helper()

	if val {
		t.Errorf("failed asserting false\n")
	}
}

// Eq is a helper function to test the equality of two passed values.
func Eq[T comparable](t *testing.T, val T, expected T) {
	t.Helper()

	if val != expected {
		t.Errorf("failed asserting '%v' == '%v' (expected)\n", val, expected)
	}
}

// Gt is a helper function to test if the passed value is greater than the threshold.
func Gt[T constraints.Ordered](t *testing.T, val T, threshold T) {
	t.Helper()

	if val <= threshold {
		t.Errorf("failed asserting %v > %v (threshold)\n", val, threshold)
	}
}

// Gt is a helper function to test if the passed value is greater than or equal to the threshold.
func Gte[T constraints.Ordered](t *testing.T, val T, threshold T) {
	t.Helper()

	if val < threshold {
		t.Errorf("failed asserting %v >= %v (threshold)\n", val, threshold)
	}
}

// Lt is a helper function to test if the passed value is less than the threshold.
func Lt[T constraints.Ordered](t *testing.T, val T, threshold T) {
	t.Helper()

	if val >= threshold {
		t.Errorf("failed asserting %v < %v (threshold)\n", val, threshold)
	}
}

// Lte is a helper function to test if the passed value is less than or equal to the threshold.
func Lte[T constraints.Ordered](t *testing.T, val T, threshold T) {
	t.Helper()

	if val > threshold {
		t.Errorf("failed asserting %v <= %v (threshold)\n", val, threshold)
	}
}

// Len is a helper function to test if the passed array is of a particular length.
func Len[T any](t *testing.T, val []T, count int) {
	t.Helper()

	if len(val) != count {
		t.Errorf("failed asserting count %v\n", count)
	}
}

// Empty is a helper function to test if the passed array is empty.
func Empty[T any](t *testing.T, val []T) {
	t.Helper()

	if len(val) != 0 {
		t.Errorf("failed asserting empty\n")
	}
}

// Match is a helper function to test if the passed string matches a regex.
func Match(t *testing.T, val string, pattern string) {
	t.Helper()

	r := regexp.MustCompile(pattern)

	if !r.MatchString(val) {
		t.Errorf("failed asserting %v matches %v (pattern)\n", val, pattern)
	}
}

// Approx is a helper function to test if the passed floating point numbers are
// approximately equal.
func Approx[T constraints.Float](t *testing.T, val T, expected T, epsilon T) {
	t.Helper()

	var approx bool

	switch any(val).(type) {
	case float32:
		approx = approx32(float32(val), float32(expected), float32(epsilon))

	default:
		approx = approx64(float64(val), float64(expected), float64(epsilon))
	}

	if !approx {
		t.Errorf("failed asserting %v approximately equals %v (expected)\n", val, expected)
	}
}

// Error is a helper function to test if an error matches a string.
func Error(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Errorf("failed asserting error is not nil\n")
	}

	Eq(t, err.Error(), expected)
}
