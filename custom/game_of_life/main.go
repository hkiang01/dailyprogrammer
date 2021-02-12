package main

import (
	"fmt"
	"log"
	"math/rand"
)

const length = 10
const width = 10

type grid struct {
	cells [length][width]bool
}

func newGrid() grid {
	cells := [length][width]bool{}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if rand.Intn(2) == 0 {
				cells[i][j] = false
			} else {
				cells[i][j] = true
			}
		}
	}
	return grid{cells}
}

func (grid *grid) printGrid() {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid.cells[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (grid *grid) iterate() {
	newGrid := newGrid()
	// (i,j) are the coordinates of the old
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			aliveNeighbors := 0
			// (k,l) are the neighbors of (i,j), within 1 row and 1 col
			for k := i - 1; k >= 0 && k < length && k <= i+1; k++ {
				for l := j - 1; l >= 0 && l < width && l <= j+1; l++ {
					// exclude self
					if k == i && l == j {
						continue
					}
					// exclude diagonals
					// if (k == i-1 || k == i+1) && (l == j-1 || l == j+1) {
					// 	continue
					// }
					if grid.cells[k][l] {
						aliveNeighbors++
					}
				}
			}
			if grid.cells[i][j] {
				if aliveNeighbors < 2 {
					newGrid.cells[i][j] = false
				} else if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newGrid.cells[i][j] = true
				} else if aliveNeighbors > 3 {
					newGrid.cells[i][j] = false
				} else {
					log.Fatal("Invalid number of alive neighbors: %d", aliveNeighbors)
				}
			} else {
				if aliveNeighbors == 3 {
					newGrid.cells[i][j] = true
				} else {
					newGrid.cells[i][j] = false
				}
			}
		}
	}

	grid.cells = newGrid.cells
}

func main() {
	cells := [length][width]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, true, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, true, true, false, false, false, false, false},
		{false, false, true, true, false, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false}}
	grid := grid{cells}
	fmt.Println("Grid created")
	grid.printGrid()

	fmt.Println("Iterating...")
	grid.iterate()
	grid.printGrid()
}
