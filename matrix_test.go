package matrix

import (
	"testing"
)

func TestNew(t *testing.T) {
	m1 := Matrix{
		2, 3,
		[]Row{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	m1copy := m1.Copy()

	if !m1.Equals(m1copy) {
		t.Fatalf("matrix.New isn't making an identical copy")
	}

	m1copy.Rows[1][1] = -5

	if m1.Equals(m1copy) {
		t.Fatalf("matrix.New modifying the copy modifies the original. Probably didn't make an actual copy, but are referencing the original Matrix' values")
	}
}

func TestNewZero(t *testing.T) {
	m23 := NewZero(2, 3)

	expected := Matrix{
		2, 3,
		[]Row{
			{0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0},
		},
	}

	if numRows, numCols := m23.Size(); numRows != 2 || numCols != 3 {
		t.Fatalf("Matrix size is not reported correctly")
	}

	if !m23.Equals(expected) {
		t.Fatalf("Matrix Equals is broken")
	}
}

func TestGetRow(t *testing.T) {
	m34 := Matrix{
		3, 4,
		[]Row{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
		},
	}

	expected := Row{21, 22, 23, 24}

	if row, err := m34.GetRow(1); err != nil || !row.Equals(expected) {
		t.Fatalf("GetRow is broken")
	}

	if _, err := m34.GetRow(100); err == nil {
		t.Fatalf("GetRow isn't checking bounds properly")
	}
}

func TestGetColumn(t *testing.T) {
	m34 := Matrix{
		3, 4,
		[]Row{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
			{31, 32, 33, 34},
		},
	}

	expected := Column{13, 23, 33}

	if column, err := m34.GetColumn(2); err != nil || !column.Equals(expected) {
		t.Fatalf("GetColumn is broken")
	}

	if _, err := m34.GetColumn(10); err == nil {
		t.Fatalf("GetColumn isn't checking bounds properly")
	}
}

func TestGetTranspose(t *testing.T) {
	m24 := Matrix{
		2, 4,
		[]Row{
			{11, 12, 13, 14},
			{21, 22, 23, 24},
		},
	}

	expected := Matrix{
		4, 2,
		[]Row{
			{11, 21},
			{12, 22},
			{13, 23},
			{14, 24},
		},
	}

	if !m24.GetTranspose().Equals(expected) {
		t.Fatalf("GetTranspose is broken")
	}
}

func TestSwapRows(t *testing.T) {
	m1, _ := New([]Row{
		{1, 1},
		{2, 2},
		{3, 3},
	})

	expected, _ := New([]Row{
		{3, 3},
		{2, 2},
		{1, 1},
	})

	m1.SwapRows(0, 2)
	if !m1.Equals(expected) {
		t.Fatalf("Swap Rows is broken")
	}

	if err := m1.SwapRows(0, 100); err == nil {
		t.Fatalf("SwapRows shouldn't work with row indexes outside row bounds")
	}
}

func TestGaussianReduce(t *testing.T) {
	m1, _ := New([]Row{
		{1.0, 2.5, -3.0},
	})
	expected1, _ := New([]Row{
		{1.0, 2.5, -3.0},
	})
	if !m1.GaussianReduce().Equals(expected1) {
		t.Errorf("m1.GaussianReduce is broken")
	}

	m2, _ := New([]Row{
		{2.0, -4.0, 1.5},
		{6.0, -4.0, 0.0},
	})
	expected2, _ := New([]Row{
		{1.0, -2.0, 0.75},
		{0.0, 1.0, -0.5625},
	})
	if !m2.GaussianReduce().Equals(expected2) {
		t.Errorf("m2.GaussianReduce is broken")
	}

	m3, _ := New([]Row{
		{1.0, 2.0, 3.0},
		{1.0, 2.0, 3.0},
		{2.0, 2.0, 4.0},
	})
	expected3, _ := New([]Row{
		{1.0, 2.0, 3.0},
		{0.0, 1.0, 1.0},
		{0.0, 0.0, 0.0},
	})
	if !m3.GaussianReduce().Equals(expected3) {
		t.Errorf("m3.GaussianReduce is broken")
	}
}
