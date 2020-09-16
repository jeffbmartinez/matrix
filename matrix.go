package matrix

import (
	"errors"
	"fmt"
)

// Matrix is a rectangular matrix (can't have rows of different lengths)
// For all matrix operations, rows and columns are zero-indexed, meaning
// the first row has index 0, the third column has index 2, and so on.
type Matrix struct {
	NRows int
	NCols int
	Rows  []Row
}

// New creates a new matrix from the rows.
// Note that if rows is empty (len(row) == 0) then the Matrix will be created with row and column length 0
// To create a Matrix with zero rows and non-zero columns, use `NewZero(0, numColumns)`` instead
func New(rows []Row) (Matrix, error) {
	nRows := len(rows)

	if nRows == 0 {
		return Matrix{0, 0, rows}, nil
	}

	rowsCopy := []Row{}

	nCols := len(rows[0])
	for _, row := range rows {
		if len(row) != nCols {
			return Matrix{}, errors.New("Matrix does not support jagged matrices. All rows must be of the same length")
		}

		rowsCopy = append(rowsCopy, row.Copy())
	}

	return Matrix{nRows, nCols, rowsCopy}, nil
}

func NewZero(nRows int, nCols int) Matrix {
	rows := []Row{}

	for i := 0; i < nRows; i++ {
		zeroRow := make(Row, nCols)
		rows = append(rows, zeroRow)
	}

	return Matrix{nRows, nCols, rows}
}

func (m Matrix) Copy() Matrix {
	// Can safely ignore error since m is already guaranteed to be a well defined Matrix
	matrixCopy, _ := New(m.Rows)
	return matrixCopy
}

// Size returns (number of rows, number of columns)
func (m Matrix) Size() (int, int) {
	return m.NRows, m.NCols
}

func (m Matrix) Equals(other Matrix) bool {
	mRows, mCols := m.Size()
	otherRows, otherCols := other.Size()
	if mRows != otherRows || mCols != otherCols {
		return false
	}

	for i, row := range m.Rows {
		if !row.Equals(other.Rows[i]) {
			return false
		}
	}

	return true
}

func (m Matrix) GetRow(row int) (Row, error) {
	if row >= m.NRows {
		return Row{}, errors.New("The row index argument is outside of the matrix' bounds")
	}

	return m.Rows[row], nil
}

func (m Matrix) GetColumn(col int) (Column, error) {
	if col >= m.NCols {
		return Column{}, errors.New("The column inddex argument is outside of the matrix' bounds")
	}

	column := Column{}
	for _, row := range m.Rows {
		column = append(column, row[col])
	}

	return column, nil
}

// Get is equivalent to `m.Rows[row][col]`` but will do bounds checking
func (m Matrix) Get(row int, col int) (float64, error) {
	if row >= m.NRows || col >= m.NCols {
		return 0.0, errors.New("You can retrieve a value outside the row/column bounds of the matrix")
	}

	return m.Rows[row][col], nil
}

// Set is equivalent to `m.Rows[row][col] = value` but will do bounds checking
func (m *Matrix) Set(value float64, row int, col int) error {
	if row >= m.NRows || col >= m.NCols {
		return errors.New("You can't set a value outside the row/column bounds of the matrix")
	}

	m.Rows[row][col] = value
	return nil
}

func (m Matrix) GetTranspose() Matrix {
	columns := []Column{}

	for i := 0; i < m.NCols; i++ {
		column, _ := m.GetColumn(i)
		columns = append(columns, column)
	}

	// Can safely ignore err because it's impossible to accidentally create a jagged matrix here
	transposedMatrix, _ := New(columns)
	return transposedMatrix
}

func (m *Matrix) SwapRows(r1 int, r2 int) error {
	if r1 >= m.NRows || r2 >= m.NRows {
		return errors.New("One or both row indexes are outside the bounds of this matrix")
	}

	if r1 == r2 {
		return nil
	}

	m.Rows[r1], m.Rows[r2] = m.Rows[r2], m.Rows[r1]
	return nil
}

func (m Matrix) GaussianReduce() Matrix {
	reducedMatrix := m.Copy()

	if m.NRows == 0 || m.NCols == 0 {
		return reducedMatrix
	}

	currentColumn := 0
	for r := range reducedMatrix.Rows {

		// find first row with a useful pivot (element != 0)
		col, _ := reducedMatrix.GetColumn(currentColumn)

		pivotRow := -1
		for i, element := range col[r:] {
			if !Equals(element, 0.0) {
				pivotRow = r + i
				break
			}
		}

		if pivotRow == -1 {
			currentColumn += 1
			continue
		}

		// Swap that element with the first element (if it isn't the first already)
		reducedMatrix.SwapRows(r, pivotRow)

		// Make the pivot equal to 1.0 (Divide all elements of that row by the pivot element)
		reducedMatrix.Rows[r] = reducedMatrix.Rows[r].ScalarMult(1.0 / reducedMatrix.Rows[r][currentColumn])

		// for each additional row:
		//   Calculate the multiple of the pivot to add to the next row to make it zero
		//   Add that multiple of that row to the other row
		for r2 := range reducedMatrix.Rows[r+1:] {
			negationFactor := -reducedMatrix.Rows[r+r2+1][currentColumn] / reducedMatrix.Rows[r][currentColumn]
			negationRow := reducedMatrix.Rows[r].ScalarMult(negationFactor)
			reducedMatrix.Rows[r+r2+1], _ = reducedMatrix.Rows[r+r2+1].Add(negationRow)
		}

		currentColumn += 1
	}

	fmt.Println(reducedMatrix)

	return reducedMatrix
}
