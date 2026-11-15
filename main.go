package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/seanhammell/abacus/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Calculator REPL!")
	fmt.Println("-------------------------------")

	for {
		fmt.Print(">> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		if strings.ToLower(input) == "q" {
			break
		}

		lexer := calculator.NewLexer(input)
		for token := lexer.NextToken(); token.Type != calculator.EOF; token = lexer.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}
