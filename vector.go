package matrix

/*
Vector is a math vector.
It uses float64 as the basic scalar type.
DefaultEpsilon is the epsilon value for comparing values.
It correctly handles comparisons and math operations with math.Inf() and NaN
*/

import (
	"errors"
	"math"
)

type Vector []float64
type Row = Vector
type Column = Vector

func (v Vector) Copy() Vector {
	newVector := make(Vector, len(v))
	copy(newVector, v)
	return newVector
}

func (v Vector) ScalarMult(c float64) Vector {
	newVector := v.Copy()

	for i := range newVector {
		newVector[i] *= c
	}

	return newVector
}

// Equals compares the Vector to another Vector and returns true if they are the same length
// and respective elements are no more than DefaultEpsilon apart in difference.
func (v Vector) Equals(other Vector) bool {
	if len(v) != len(other) {
		return false
	}

	for i, element := range v {
		if !Equals(element, other[i]) {
			return false
		}
	}

	return true
}

// Magnitude calculates the magnitude of the vector
// An empty vector (with no elements) will return a magnitude of zero
func (v Vector) Magnitude() float64 {
	sumOfSquares := 0.0

	for _, element := range v {
		sumOfSquares += element * element
	}

	return math.Sqrt(sumOfSquares)
}

func (v Vector) Add(other Vector) (Vector, error) {
	if len(v) != len(other) {
		return Vector{}, errors.New("Cannot add vectors of different length")
	}

	result := v.Copy()

	for i, element := range other {
		result[i] += element
	}

	return result, nil
}

func (v Vector) Dot(other Vector) (float64, error) {
	if len(v) != len(other) {
		return 0.0, errors.New("Cannot take dot product of vectors of different length")
	}

	sum := 0.0

	for i, element := range v {
		sum += element * other[i]
	}

	return sum, nil
}
