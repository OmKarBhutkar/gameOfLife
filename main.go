package main

import (
	"fmt"
	"math/rand"
)

var (
	col         int = 9
	row         int = 9
	grid        [][]int
	generations int = 9
)

func main() {
	// Create 2d grid as per the given inputs
	grid = create2DGrid()
	//Fill grid with random data
	fillGrid()
	printGrid()

	//To Number of generations
	for n := 0; n < generations; n++ {
		regeneration()
		fmt.Println("-----------")
		printGrid()
	}

}

// Create 2d grid as per the given inputs
func create2DGrid() [][]int {
	grid := make([][]int, col)
	for i := range grid {
		grid[i] = make([]int, row)
	}
	return grid
}

// Fill grid with random data
func fillGrid() {
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			grid[i][j] = (rand.Intn(99) % 2)
		}
	}
}

// print data of given grid
func printGrid() {
	for _, col := range grid {
		for _, row := range col {
			if row == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// create next generation
func regeneration() {
	next := create2DGrid()
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			neighbors := numberOfNeighbors(i, j)
			currentState := grid[i][j]
			if currentState == 0 && neighbors == 3 {
				next[i][j] = 1
			} else if currentState == 1 && (neighbors < 2 || neighbors > 3) {
				next[i][j] = 0
			} else {
				next[i][j] = currentState
			}
		}
	}
	grid = next
}

// Calculate number of neighbors for each cell.
func numberOfNeighbors(i, j int) int {
	count := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			lcol := (i + x)
			lrow := (j + y)
			if lcol < 0 || lrow < 0 || lcol >= col || lrow >= row {
				continue
			}

			count += grid[lcol][lrow]
		}
	}
	count -= grid[i][j]
	return count
}
