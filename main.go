package main

import (
	"fmt"
	"gdejong/interpreter-in-go/pkg/repl"
	"os"
)

func main() {
	fmt.Println("Hello! This is the Monkey programming language!")
	fmt.Println("Type in some Monkey commands and press enter to see the lexer in action.")

	repl.Start(os.Stdin, os.Stdout)
}
