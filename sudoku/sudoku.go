package sudoku

import (
	"fmt"
)

type Sudoku struct {
	I    int     `json:"i"` // index of current row (algorithm purpose)
	J    int     `json:"j"` // index of current column (algorithm purpose)
	Size int     `json:"size"`
	Grid [][]int `json:"grid"`
}

var sudokuZero Sudoku

func InitGrid() Sudoku {
	// grid := [][]int{
	// 	[]int{1, 0, 0, 4},
	// 	[]int{0, 2, 0, 0},
	// 	[]int{0, 0, 0, 0},
	// 	[]int{0, 0, 0, 0},
	// }

	grid := [][]int{
		[]int{1, 0, 0, 4},
		[]int{0, 0, 2, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 3, 0},
	}

	sudoku := Sudoku{
		I:    0,
		J:    0,
		Size: 4,
		Grid: grid,
	}

	saveSudokuZero(sudoku)
	return sudoku
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

func (m *Sudoku) ChangePossibleValue(i int, j int) int {
	value := 0
	curValue := m.GetValue(i, j)
	if curValue > m.Size {
		return 0
	}

	for d := curValue + 1; d <= m.Size; d++ {
		if m.isPossible(i, j, d) && m.isNull(i, j) {
			return d
		}
	}

	return value
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

	return values.Distinct()
}
