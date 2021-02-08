package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var inputPattern = regexp.MustCompile(`(?P<a>\w+) (?P<op>[+\-*\/]+) (?P<b>\w+)`)

func parseInput() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	matches := inputPattern.FindStringSubmatch(line)
	if len(matches) < 3 {
		log.Fatal("Invalid input")
	}
	a := matches[1]
	op := matches[2]
	b := matches[3]
	return a, op, b
}

func main() {
	a, op, b := parseInput()
	fmt.Printf("%s %s %s\n", a, op, b)
}
