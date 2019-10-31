package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := "non"

	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	fmt.Printf("your input is %s!\n", input)

	if input == "error" {
		log.Fatal("error raise")
	}
}
