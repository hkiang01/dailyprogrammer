package main

import (
	"fmt"
	"log"
	"strconv"
)

func prompt(prompt string) string {
	fmt.Println(prompt)
	var field string
	_, err := fmt.Scanln(&field)
	if err != nil {
		log.Fatal(err)
	}
	return field
}

func main() {
	name := prompt("What is your name?")
	ageStr := prompt("How old are you?")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		log.Fatal(err)
	}
	user := prompt("What is your username?")
	fmt.Println(fmt.Sprintf("Your name is %s, you are %d years old, and your username is %s", name, age, user))
}
