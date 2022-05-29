package main

import (
	"fmt"
	"os"
)

// Modify the echo program to print the index and value of each of its arguments, one per line.
func main() {
	for index, val := range os.Args[1:] {
		s := "index: " + fmt.Sprintf("%d", index) + ", value: " + val
		fmt.Println(s)
	}
}
