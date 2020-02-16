package sudoku

import (
	"fmt"
	"math"
)

type Sudoku struct {
	I            int     `json:"i"` // index of current row (algorithm purpose)
	J            int     `json:"j"` // index of current column (algorithm purpose)
	Size         int     `json:"size"`
	Grid         [][]int `json:"grid"`
	IsSolvable   bool    `json:"is_solvable"`
	IsStepByStep bool    `json:"is_step_by_step"`
}

var sudokuZero Sudoku

func Build(sdk Sudoku) Sudoku {
	grid := initiateGrid(sdk.IsSolvable)

	sdk = Sudoku{
		I:          0,
		J:          0,
		Size:       9, // it's must be quadratic integer
		Grid:       grid,
		IsSolvable: sdk.IsSolvable,
	}

	saveSudokuZero(sdk)
	return sdk
}

func (m *Sudoku) PrintGrid() {
	for _, row := range m.Grid {
		fmt.Println(row)
	}
}

func (m *Sudoku) PrintSudokuZero() {
	for _, row := range sudokuZero.Grid {
		fmt.Println(row)
	}
}

func (m *Sudoku) InitIndex() bool {
	if !m.isNull(0, 0) {
		return m.NextIndex()
	}
	return true
}

func (m *Sudoku) NextIndex() bool {
	if m.J < m.Size-1 {
		m.J += 1
	} else if m.I < m.Size-1 {
		m.I += 1
		m.J = 0
	} else {
		return false
	}

	if !sudokuZero.isNull(m.I, m.J) {
		return m.NextIndex()
	}

	return true
}

func (m *Sudoku) PrevIndex() bool {
	if m.J > 0 {
		m.J -= 1
	} else if m.I > 0 {
		m.I -= 1
		m.J = m.Size - 1
	} else {
		return false
	}

	if !sudokuZero.isNull(m.I, m.J) {
		return m.PrevIndex()
	}

	return true
}

func (m *Sudoku) PickPossibleValue(i int, j int, minValue int) int {
	for d := minValue; d <= m.Size; d++ {
		if m.isPossible(i, j, d) {
			return d
		}
	}

	return 0
}

func (m *Sudoku) SetValue(i int, j int, value int) {
	m.Grid[i][j] = value
}

func (m *Sudoku) GetValue(i int, j int) int {
	return m.Grid[i][j]
}

// private method
func saveSudokuZero(sdk Sudoku) {
	gridZero := make([][]int, len(sdk.Grid))
	for i := range sdk.Grid {
		gridZero[i] = make([]int, len(sdk.Grid[i]))
		copy(gridZero[i], sdk.Grid[i])
	}

	sudokuZero = Sudoku{
		I:    sdk.I,
		J:    sdk.J,
		Size: sdk.Size,
		Grid: gridZero,
	}
}

func (m *Sudoku) isPossible(i int, j int, d int) bool {
	return !m.notPossibleValues(i, j).Contains(d)
}

func (m *Sudoku) isNull(i int, j int) bool {
	if m.GetValue(i, j) != 0 {
		return false
	}
	return true
}

func (m *Sudoku) notPossibleValues(i int, j int) Array {
	var values Array

	for idx := 0; idx < m.Size; idx++ {
		rowValue := m.GetValue(i, idx)
		if rowValue != 0 {
			values = append(values, rowValue)
		}
	}

	for idx := 0; idx < m.Size; idx++ {
		clmValue := m.GetValue(idx, j)
		if clmValue != 0 {
			values = append(values, clmValue)
		}
	}

	dim := int(math.Sqrt(float64(m.Size)))

	ii := (i / dim) * dim
	jj := (j / dim) * dim

	for k := ii; k < ii+dim; k++ {
		for l := jj; l < jj+dim; l++ {
			boxValue := m.GetValue(k, l)
			if boxValue != 0 {
				values = append(values, boxValue)
			}
		}
	}

	return values.Distinct()
}

func initiateGrid(isSolvable bool) [][]int {
	grid := [][]int{}

	if isSolvable {
		grid = [][]int{
			[]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
			[]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
			[]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
			[]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
			[]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
			[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
			[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
			[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
			[]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}
	} else {
		grid = [][]int{
			[]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
			[]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
			[]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
			[]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
			[]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
			[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
			[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
			[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
			[]int{7, 0, 0, 0, 8, 0, 0, 7, 9},
		}
	}
	return grid
}
