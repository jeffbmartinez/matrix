package matrix

import (
	"math"
	"testing"
)

func TestScalarMult(t *testing.T) {
	negativeInfinity := math.Inf(-1)
	positiveInfinity := math.Inf(1)
	undefined := math.NaN()

	row := Row{negativeInfinity, -1.3, 0.0, 1.3, positiveInfinity}

	scalars := []float64{
		negativeInfinity,
		-2.1,
		0.0,
		2.1,
		positiveInfinity,
	}

	expected := []Row{
		Row{positiveInfinity, positiveInfinity, undefined, negativeInfinity, negativeInfinity},
		Row{positiveInfinity, 2.73, 0.0, -2.73, negativeInfinity},
		Row{undefined, 0.0, 0.0, 0.0, undefined},
		Row{negativeInfinity, -2.73, 0.0, 2.73, positiveInfinity},
		Row{negativeInfinity, negativeInfinity, undefined, positiveInfinity, positiveInfinity},
	}

	for i, scalar := range scalars {
		result := row.ScalarMult(scalar)

		if !result.Equals(expected[i]) {
			t.Fatalf("Failed with argument %v\nexpected: %v\nactual: %v", scalars[i], expected[i], result)
		}
	}
}
