package main

import (
	"fmt"
	"strings"
)

// signature
type StringCallback func(string)

func processSlice(items []string, callback StringCallback) {
	fmt.Println("Processing...")
	for _, item := range items {
		callback(item)
	}
	fmt.Println("Processed.")
}

func main() {
	names := []string{"John", "Doe", "Alice"}

	uppercaseCallback := func(s string) {
		fmt.Printf(" - In uppercase %s\n", strings.ToUpper(s))
	}

	processSlice(names, uppercaseCallback)
}
