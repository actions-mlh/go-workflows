package parser

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	dir, err := os.Open("./yaml/")
	if err != nil {
		t.Fatal("error reading dir yaml")
	}
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		t.Fatal("error reading files from yaml")
	}

	for _, file := range files {
		data, err := os.ReadFile("./yaml/" + file)
		if err != nil {
			t.Fatalf("error reading test file %v",file)
		}
		_, err = Parse(data)
		if err != nil {
			t.Errorf("failed to read file %v", file)
		}
	}
}
