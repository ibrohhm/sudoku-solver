package sudoku

type Array []int

func (arr Array) Length() int {
	return len(arr)
}

func (arr Array) Contains(d int) bool {
	for _, value := range arr {
		if value == d {
			return true
		}
	}
	return false
}

func (arr Array) Distinct() Array {
	var result Array
	for _, value := range arr {
		if !result.Contains(value) {
			result = append(result, value)
		}
	}
	return result
}
