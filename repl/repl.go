package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		fmt.Println(input)
	}
}
