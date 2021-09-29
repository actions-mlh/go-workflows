package main

import (
	"fmt"
	"os"
	"strings"
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
	errorMsg := yaml.Unmarshal(data, obj)
	if errorMsg == nil {
		fmt.Println("yaml parse ok! object:\n")
		fmt.Println(obj)
	} else {
		// split yaml file into lines
		lines := strings.Split(string(data), "\n")

		// get line number of error
		lineNoString := strings.Fields(errorMsg.Error())[2]
		lineNoString = lineNoString[:len(lineNoString)-1]
		var lineNo int
		_, err := fmt.Sscan(lineNoString, &lineNo)
		if err != nil {
			panic("line no. not parsed correctly")
		}
		
		fmt.Println(errorMsg)
		fmt.Println(lines[lineNo])
		os.Exit(1)
	}
}
