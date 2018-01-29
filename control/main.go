package main

import (
	"fmt"
)

func main() {
	// if
	a := 5
	if a > 10 {
		fmt.Printf("a > 10\n")
	} else if a > 20 {
		fmt.Printf("a > 20\n")
	} else {
		fmt.Printf("else\n")
	}
}
