package main

import (
	"fmt"
	"os"
)

// report_generator - Generate data reports
func report_generator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Report-Generator")
	fmt.Println("  Generate data reports")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	report_generator(path)
}
