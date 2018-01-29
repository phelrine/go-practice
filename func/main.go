package main

import (
	"fmt"
)

func main() {
	fmt.Printf("calc(2, 3) = %d\n", calc(2, 3))
	quo, remain := calcQuoAndRemain(10, 3)
	fmt.Printf("quo = %d, remain = %d\n", quo, remain)

	// 関数を変数に代入して利用することができる
	f1 := func(x int) int {
		return 5 * x
	}
	fmt.Printf("f1(1) = %d\n", f1(1))
	fmt.Printf("f1(3) = %d\n", f1(3))
}

func calc(a int, b int) int {
	return a + b
}

func calcQuoAndRemain(a, b int) (int, int) {
	return a / b, a % b
}
