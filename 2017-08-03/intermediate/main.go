package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type cell struct {
	y int
	x int
}

var maze [][]string
var visited [][]bool
var path []cell

func parse(filePath string) (cell, error) {
	// get lines of file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")

	// get starting cell
	startLine := lines[0]
	re := regexp.MustCompile(`\((?P<x>\d),(?P<x>\d)\)`)
	matches := re.FindStringSubmatch(startLine)
	x, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal(err)
	}
	startCell := cell{y: y, x: x}

	// parse maze
	for _, line := range lines[1:] {
		rawCells := strings.Split(line, " ")
		mazeLine := make([]string, 0)
		for _, rawCell := range rawCells {
			if rawCell != "" {
				mazeLine = append(mazeLine, rawCell)
			}
		}
		maze = append(maze, mazeLine)
	}
	if len(maze) == 0 {
		return cell{y: 0, x: 0}, errors.New("empty maze")
	}
	numCols := len(maze[0])
	for _, row := range maze {
		if len(row) != numCols {
			return cell{y: 0, x: 0}, errors.New("maze has uneven number of rows")
		}
	}
	return startCell, nil
}

func initializeVisited(nRows int, nCol int) {
	for i := 0; i < nRows; i++ {
		visited = append(visited, make([]bool, nCol))
	}
}

func getDirection(s string) (cell, error) {
	switch s {
	case "h":
		return cell{x: 0, y: 0}, nil
	case "n":
		return cell{y: -1, x: 0}, nil
	case "ne":
		return cell{y: -1, x: 1}, nil
	case "e":
		return cell{y: 0, x: 1}, nil
	case "se":
		return cell{y: 1, x: 1}, nil
	case "s":
		return cell{y: 1, x: 0}, nil
	case "sw":
		return cell{y: 1, x: -1}, nil
	case "w":
		return cell{y: 0, x: -1}, nil
	case "nw":
		return cell{y: -1, x: -1}, nil
	default:
		return cell{y: 0, x: 0}, fmt.Errorf("invalid param %s", s)
	}
}

func iterate(curr cell, direction cell) cell {
	return cell{y: curr.y + direction.y, x: curr.x + direction.x}
}

func outOfBounds(c cell) bool {
	return c.y < 0 || c.x < 0 || c.y >= len(maze) || c.x >= len(maze)
}

func walk(direction cell, curr cell, depth int) bool {
	// end condition: are we home?
	if maze[curr.y][curr.x] == "h" {
		path = append(path, curr)
		return true
	}
	if visited[curr.y][curr.x] {
		return false
	}
	visited[curr.y][curr.x] = true
	if depth > 0 {
		// follow current direction
		nextFollowingDirection := iterate(curr, direction)
		if !outOfBounds(nextFollowingDirection) {
			if walk(direction, nextFollowingDirection, depth-1) {
				path = append(path, curr)
				return true
			}
		}

		// follow arrow
		arrow := maze[curr.y][curr.x]
		arrowDirection, err := getDirection(arrow)
		if err != nil {
			log.Fatal(err)
		}
		nextFollowingArrow := iterate(curr, arrowDirection)
		if !outOfBounds(nextFollowingArrow) {
			if walk(arrowDirection, nextFollowingArrow, depth-1) {
				path = append(path, curr)
				return true
			}
		}
	}
	return false
}

func main() {
	start, err := parse("maze.txt")
	initializeVisited(len(maze), len(maze[0]))
	if err != nil {
		log.Fatal(err)
	}
	direction, err := getDirection(maze[start.y][start.x])
	if err != nil {
		log.Fatal(err)
	}
	routeFound := walk(direction, start, 20)
	if routeFound {
		fmt.Println(path)
	} else {
		fmt.Println("No route found")
	}
}
