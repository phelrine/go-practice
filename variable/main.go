package main

import (
	"fmt"
)

func main() {
	// 宣言と代入
	var a int
	a = 1
	fmt.Printf("int a= %d\n", a)

	// 宣言と初期化を同時に行う
	b := 2
	fmt.Printf("int b = %d\n", b)

	c := "文字列"
	fmt.Printf("string c = %s\n", c)

	d := 1.0
	fmt.Printf("float d = %0.3f\n", d)

	type myInt int
	var e myInt
	e = 1
	fmt.Printf("myInt e = %d\n", e)
}
