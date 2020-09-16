package matrix

import "math"

func Equals(n1 float64, n2 float64) bool {
	// DefaultEpsilon is used for comparisons of scalar values.
	// 1e-10 because 2^-32 ~= 2.33e-10 seems reasonable for a float64 epsilon
	const epsilon float64 = 1e-10

	return EqualsWithEpsilon(n1, n2, epsilon)
}

func EqualsWithEpsilon(n1 float64, n2 float64, epsilon float64) bool {
	if math.IsNaN(n1) {
		if math.IsNaN(n2) {
			return true
		} else {
			return false
		}
	}

	// n1 is a number at this point
	if math.IsNaN(n2) {
		return false
	}

	diff := math.Abs(n1 - n2)
	if diff > epsilon {
		return false
	}

	return true
}
