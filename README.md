# sudoku-solver
This sudoku solver using backtracking algorithm

Backtracking is brute force search technique to find solution in computational problem and find the solution of sudoku is one of them. I don't wanna talk too much about backtracking, so lets talk about 
the implementation in this sudoku-solver

## The Algorithm
Sortly, find the possible value of each cell of the soduku grid from left to the right and from top to the below, if the cell has no possible value go back to previous cell and change it's value

1. pick the first empty cell (left --> right, top --> bellow)
2. find the possible minimum value of the cell (start from 1 to 9)
3. if it has possible value, move to next cell and repeat step 2
4. if it has no possible value, back to previous cell and repeat step 2 with minimum value higher than it's current value
5. do the steps until each cell has value or the grid actually has no solution at all 

when each cell has possible value it's actually the solution for the sudoku grid and you can end the computation, but how do we know to state that the soduku has no solution?. 
It's related with `step 4` when the cell has no possible value, back to previous cell and find the possible value of the cell with minimum value must be higher than it's current value (ie: `current value: 1` so range `possible value` is from 2 to 9), but when the cell still has no possible value, you need to go back to the previous cell and remove the `current value` cell (`nil` value), if this bad state continous and reach the first cell so this soduku has no solution.

## The Implementation

the sudoku grid can be represent as array with dimension `n X n` with `n` is it size
in this repo, I create model for sudoku like this

```
type Sudoku struct {
	I            int     `json:"i"` // index of current row (algorithm purpose)
	J            int     `json:"j"` // index of current column (algorithm purpose)
	Size         int     `json:"size"`
	Grid         [][]int `json:"grid"`
	IsSolvable   bool    `json:"is_solvable"`
	IsStepByStep bool    `json:"is_step_by_step"`
}
```

I create two example of initial sudoku grid, one is solvable and another one is unsolvable to see the different, it's render by `IsSolvable` value. Also I always print each iteration step to see the progress but if you wanna check one by one of the step, make it `IsStepByStep` has `true` value.

The first cell is cell with index (0,0), if it has value check next cell (0,1) and so on, or you can see this function to initiate first empty cell

```
func (m *Sudoku) InitIndex() bool {
	...
	return true
}
```


sudoku has roles
1. each cell of any row has distinct value
2. each cell of any column has distinct value
3. each cell of any box has distinct value
4. possible value of each cell is from 1 to 9

it's easy to find all impossible values first and then find the minimum value from the rest of the values, check these functions

```
func (m *Sudoku) PickPossibleValue(i int, j int, minValue int) int {
  ...
}

func (m *Sudoku) isPossible(i int, j int, d int) bool {
	...
}

func (m *Sudoku) notPossibleValues(i int, j int) Array {
 ...
}

```

if it has possible value do `step 3` and if it's not do `step 4` and loop the iteration until find all solution or not at all, I am to lazy to repeat describe how to end this iteration, please read the description above or direct read the code, actually I made it readable.

## Conclution

Backtracking is brute-force so please don't use it when time complexity is realy matter for you, I just implement it from this [wiki](https://en.wikipedia.org/wiki/Sudoku_solving_algorithms#Backtracking). Fell free to change the initate sudoku grid or change it's size, but make sure that is has to be `quadratic integer` or you can modify the algorithm so that it's can find the solution of soduku with `n X m` dimension, whatever.

## Run the programm

```
go run main.go
```
