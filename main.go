package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello! It's a me, al-gore")
	fmt.Printf("=========================\n\n")
	fmt.Printf("Running sort algorithms\n\n")

	tests := generateSortTests()
	data := generateDataSet()
	runTests(tests, data)

	fmt.Println("All done, bye bye!")
}
