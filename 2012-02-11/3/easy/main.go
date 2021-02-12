package main

import (
	"fmt"
	"log"
)

const cipherOffset = 10

func prompt(prompt string) string {
	fmt.Println(prompt)
	var field string
	_, err := fmt.Scanln(&field)
	if err != nil {
		log.Fatal(err)
	}
	return field
}

func encrypt(input string, offset int) {
	for _, c := range input {
		fmt.Print(string(rune(int(c) + offset)))
	}
	fmt.Println()
}

func decrypt(input string, offset int) {
	for _, c := range input {
		fmt.Print(string(rune(int(c) - offset)))
	}
	fmt.Println()
}

func main() {
	input := prompt("string to encrypt")
	encrypt(input, cipherOffset)
	input = prompt("string to decrypt")
	decrypt(input, cipherOffset)
}
