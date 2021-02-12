package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const Length = 5
const Width = 5

type grid struct {
	cells [Length][Width]bool
}

func newGrid() grid {
	cells := [Length][Width]bool{}
	for i := 0; i < Length; i++ {
		for j := 0; j < Width; j++ {
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
	for i := 0; i < Length; i++ {
		for j := 0; j < Width; j++ {
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
	for i := 0; i < Length; i++ {
		for j := 0; j < Width; j++ {
			aliveNeighbors := 0
			for k := i - 1; k >= 0 && k < Length && k <= i+1; k++ {
				for l := j - 1; l >= 0 && l < Width && l <= j+1; l++ {
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
	rand.Seed(time.Now().UnixNano())
	grid := newGrid()
	fmt.Println("Grid created")
	grid.printGrid()

	fmt.Println("Iterating...")
	grid.iterate()
	grid.printGrid()
}
