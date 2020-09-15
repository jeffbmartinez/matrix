package matrix

import (
	"math"
)

// DefaultEpsilon is used for comparisons of scalar values.
// 1e-10 because 2^-32 ~= 2.33e-10 seems reasonable for a float64 epsilon
const DefaultEpsilon float64 = 1e-10

type Row []float64

func (r Row) Copy() Row {
	newRow := make(Row, len(r))
	copy(newRow, r)
	return newRow
}

func (r Row) ScalarMult(c float64) Row {
	newRow := r.Copy()

	for i := range newRow {
		newRow[i] *= c
	}

	return newRow
}

// Equals compares the Row to another Row and returns true if they are the same length
// and respective elements are no more than DefaultEpsilon apart in difference.
func (r Row) Equals(other Row) bool {
	if len(r) != len(other) {
		return false
	}

	for i := range r {
		diff := math.Abs(r[i] - other[i])
		if diff > DefaultEpsilon {
			return false
		}
	}

	return true
}
