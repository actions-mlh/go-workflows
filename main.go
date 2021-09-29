package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

func main() {
	fmt.Println("hello world!")

	data, err := os.ReadFile("simple.yaml")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println(data)

	obj := make(map[interface{}]interface{})
	yaml.Unmarshal(data, obj)
	fmt.Println(obj)
}
