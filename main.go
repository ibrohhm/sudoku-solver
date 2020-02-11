package main

import (
	"fmt"
	"os"
	"os/exec"

	model "github.com/learn/backtracking/sudoku"
)

func main() {
	sdk := model.InitGrid()
	sdk.PrintGrid()

	sdk.InitIndex()
	sdk = Backtracking(sdk, false)

	fmt.Scanln()
	clearConsole()

	if sdk.I == 0 {
		printFailed(sdk)
	} else {
		printSuccess(sdk)
	}
}

func Backtracking(sdk model.Sudoku, backward bool) model.Sudoku {
	var value int
	minValue := 1
	if backward {
		minValue = sdk.GetValue(sdk.I, sdk.J) + 1
	}
	fmt.Scanln()
	clearConsole()

	value = sdk.PickPossibleValue(sdk.I, sdk.J, minValue)

	if value != 0 {
		sdk.SetValue(sdk.I, sdk.J, value)
		sdk.PrintGrid()
		if !sdk.NextIndex() {
			return sdk
		}
		fmt.Println("next")
		sdk = Backtracking(sdk, false)
	} else if value == 0 && backward {
		sdk.SetValue(sdk.I, sdk.J, 0)
		sdk.PrintGrid()
		if !sdk.PrevIndex() {
			return sdk
		}
		fmt.Println("back")
		sdk = Backtracking(sdk, true)
	} else {
		sdk.PrintGrid()
		if !sdk.PrevIndex() {
			return sdk
		}
		fmt.Println("change")
		sdk = Backtracking(sdk, true)
	}

	return sdk
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func printSuccess(sdk model.Sudoku) {
	fmt.Println("---init---")
	sdk.PrintSudokuZero()

	fmt.Println()

	fmt.Println("---result---")
	sdk.PrintGrid()
}

func printFailed(sdk model.Sudoku) {
	fmt.Println("---init---")
	sdk.PrintSudokuZero()

	fmt.Println()

	fmt.Println("---no solution found---")
}
