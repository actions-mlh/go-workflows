package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
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

	obj := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, obj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("yaml parse ok! object:\n")
	fmt.Println(obj)
}
