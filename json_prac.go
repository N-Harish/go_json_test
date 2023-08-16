package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	x := map[string]any{
		"1": "1",
		"2": "2",
		"3": map[string]string{
			"x": "y",
		},
	}

	file, _ := json.MarshalIndent(x, "", " ")
	fmt.Print(string(file))
	os.WriteFile("test.json", file, os.ModeAppend)
}
