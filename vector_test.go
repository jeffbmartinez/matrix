package matrix

import (
	"math"
	"testing"
)

func TestEquals(t *testing.T) {
	vectorEmpty1 := Vector{}
	vectorEmpty2 := Vector{}

	vector0 := Vector{0.0}
	vector1 := Vector{0.0, 1.1}
	vector2 := Vector{0.0, 1.1}
	vector3 := Vector{0.0, 2.0}

	vectorNaN1 := Vector{math.NaN()}
	vectorNaN2 := Vector{math.NaN()}

	if !vectorEmpty1.Equals(vectorEmpty2) {
		t.Fatalf("Empty vector comparison failed")
	}

	if vector0.Equals(vector1) {
		t.Fatalf("Vectors of different size shouldn't be equal")
	}

	if vector1.Equals(vector3) {
		t.Fatalf("Basic vector comparison failed when they shouldn't be equal")
	}

	if !vector1.Equals(vector2) {
		t.Fatalf("Basic vector comparison failed when they should be equal")
	}

	if !vectorNaN1.Equals(vectorNaN2) {
		t.Fatalf("Comparison of vectors with NaN as both elements is broken")
	}

	if vector0.Equals(vectorNaN1) || vectorNaN1.Equals(vector0) {
		t.Fatalf("Comparison of vectors with NaN (number vs NaN) as only one of the elements is broken")
	}
}

func TestScalarMult(t *testing.T) {
	negativeInfinity := math.Inf(-1)
	positiveInfinity := math.Inf(1)
	undefined := math.NaN()

	vector := Vector{negativeInfinity, -1.3, 0.0, 1.3, positiveInfinity}

	scalars := []float64{
		negativeInfinity,
		-2.1,
		0.0,
		2.1,
		positiveInfinity,
	}

	expected := []Vector{
		{positiveInfinity, positiveInfinity, undefined, negativeInfinity, negativeInfinity},
		{positiveInfinity, 2.73, 0.0, -2.73, negativeInfinity},
		{undefined, 0.0, 0.0, 0.0, undefined},
		{negativeInfinity, -2.73, 0.0, 2.73, positiveInfinity},
		{negativeInfinity, negativeInfinity, undefined, positiveInfinity, positiveInfinity},
	}

	for i, scalar := range scalars {
		result := vector.ScalarMult(scalar)

		if !result.Equals(expected[i]) {
			t.Fatalf("Failed with argument %v\nexpected: %v\nactual: %v", scalars[i], expected[i], result)
		}
	}
}

func TestMagnitude(t *testing.T) {
	v1 := Vector{3.0, 4.0}
	expected1 := 5.0

	if v1.Magnitude() != expected1 {
		t.Fatalf("Magnitude calculation is broken")
	}
}

func TestAdd(t *testing.T) {
	v1 := Vector{1.5, -1.5}
	v2 := Vector{-0.5, 2.0}
	v3 := Vector{0.0}

	expected := Vector{1.0, 0.5}
	if result, err := v1.Add(v2); err != nil || !result.Equals(expected) {
		t.Fatalf("Vector addition is broken")
	}

	if _, err := v2.Add(v3); err == nil {
		t.Fatalf("Shouldn't be able to add together vectors of different length")
	}
}

func TestDot(t *testing.T) {
	v1 := Vector{1.0, 2.0}
	v2 := Vector{3.0, -1.0}
	v3 := Vector{0.0}

	expected := 1.0

	if dotProduct, err := v1.Dot(v2); err != nil || !Equals(dotProduct, expected) {
		t.Fatalf("Dot product calculation is broken")
	}

	if _, err := v2.Dot(v3); err == nil {
		t.Fatalf("Shouldn't be able to take dot product of different length vectors")
	}
}
