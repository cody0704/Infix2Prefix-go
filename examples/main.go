package main

import (
	"fmt"

	"github.com/cody0704/Expression-Convert/prefix"
)

func main() {
	var expression string

	fmt.Printf("Enter expression: ")
	fmt.Scanln(&expression)
	result := prefix.Infix2Prefix(expression)

	fmt.Println("Prefix:", result)
}
