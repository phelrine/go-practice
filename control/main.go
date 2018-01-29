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

	// for
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// 無限ループ
	for {
		fmt.Println("loop")
	}
}
