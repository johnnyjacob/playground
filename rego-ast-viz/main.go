package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	data, err := os.ReadFile("./test.rego")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ParseContents((string(data)))
	if err != nil {
		fmt.Println(err)
	}
}
