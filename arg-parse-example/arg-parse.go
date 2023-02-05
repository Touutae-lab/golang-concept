package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("You can use ./arg-parse --help to see command\n")
		return
	}
	fmt.Printf("hello world %v\n", args[1])
}
