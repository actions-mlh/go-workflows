package lint

import (
	"strings"
	"strconv"
	"gopkg.in/yaml.v3"
)

// Lint returns true, -1, nil if the lint passes.
// If not, it returns false, lineNo, error.
func Lint(data []byte) (bool, int, error) {
	obj := make(map[interface{}]interface{})
	err := yaml.Unmarshal(data, obj)
	
	if err == nil {
		return true, -1, nil
	} else {
		// get line number of error		
		s := strings.Fields(err.Error())[2]
		// cut off last character (a ":")
		s = s[:len(s)-1]
		// convert to int
		lineNo, err2 := strconv.Atoi(s)
		if err2 != nil {
			panic("line no. not parsed correctly")
		}
		
		return false, lineNo, err
	}

}
