package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
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

var operandPattern = regexp.MustCompile(`(?P<magnitude>\d+)(?P<unit>\w+)`)
var prefixes = map[string]int{
	"KB": int(math.Pow(1000, 1)),
	"MB": int(math.Pow(1000, 2)),
	"GB": int(math.Pow(1000, 3)),
	"TB": int(math.Pow(1000, 4)),
	"PB": int(math.Pow(1000, 5)),

	"KiB": int(math.Pow(1024, 1)),
	"MiB": int(math.Pow(1024, 2)),
	"GiB": int(math.Pow(1024, 3)),
	"TiB": int(math.Pow(1024, 4)),
	"PiB": int(math.Pow(1024, 5)),
}

func toBytes(operand string) int {
	matches := operandPattern.FindStringSubmatch(operand)
	if len(matches) < 2 {
		log.Fatalf("Invalid operand %s", operand)
	}
	magnitude, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	unit := matches[2]
	multiplicand := prefixes[unit]
	if multiplicand == 0 {
		log.Fatalf("Invalid unit %s", unit)
	}
	return magnitude * multiplicand
}

func main() {
	a, op, b := parseInput()
	aBytes := toBytes(a)
	bBytes := toBytes(b)
	var result int
	switch op {
	case "+":
		result = aBytes + bBytes
	case "-":
		result = aBytes - bBytes
	case "*":
		result = aBytes * bBytes
	case "/":
		result = aBytes / bBytes
	default:
		log.Fatal("Operand not found")
	}
	fmt.Printf("%d", result)
}
