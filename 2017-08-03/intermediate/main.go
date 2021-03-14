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

func walk(direction cell, curr cell, depth int) (cell, error) {
	// end condition: are we home?
	if maze[curr.y][curr.x] == "h" {
		return curr, nil
	}
	if depth > 0 {
		// follow current direction
		var paths []cell
		nextFollowingDirection := iterate(curr, direction)
		if !outOfBounds(nextFollowingDirection) {
			journey, err := walk(direction, nextFollowingDirection, depth-1)
			if err != nil {
				return cell{y: -1, x: -1}, fmt.Errorf("%s out of bounds", nextFollowingDirection)
			}

		} else {
			return cell{y: -1, x: -1}, fmt.Errorf("%s out of bounds", nextFollowingDirection)
		}

		// follow arrow
		arrow := maze[curr.y][curr.x]
		arrowDirection, err := getDirection(arrow)
		if err != nil {
			log.Fatal(err)
		}
		nextFollowingArrow := iterate(curr, arrowDirection)
		if !outOfBounds(nextFollowingArrow) {
			paths = append(paths, nextFollowingArrow)
		}

		var validPaths []cell
		for _, path := range paths {
			journey := walk()
		}

		if !outOfBounds(nextFollowingDirection) {
			pathsFollowingDirection := walk(direction, nextFollowingDirection, depth-1)
			if pathsFollowingDirection != nil {
				paths = append(paths, pathsFollowingDirection...)
			}
		}

		// // follow arrow

		// walk()

		return paths
	} else {
		return nil
	}
}

func main() {
	start, err := parse("maze.txt")
	if err != nil {
		log.Fatal(err)
	}
	direction, err := getDirection(maze[start.y][start.x])
	if err != nil {
		log.Fatal(err)
	}
	path := walk(direction, start, 3)
	fmt.Println(path)
}
