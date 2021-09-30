package main

import (
	"fmt"
	"os"
	
	"yaml-parser/lint"
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

	err = lint.Lint(data)

	if err == nil {
		fmt.Println("lint passed!")
	} else {
		fmt.Println("lint failed :(")
		fmt.Println(err.Error())
	}
}
