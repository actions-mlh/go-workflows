package main

import (
	"fmt"
	"os"
	
	"gh-actions-checker/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename.")
		os.Exit(1)
	}
	
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := parser.Parse(data)

	if err == nil {
		fmt.Println("lint passed!")
		fmt.Println(out)
	} else {
		fmt.Println("lint failed :(")
		fmt.Println(err.Error())
	}
}
