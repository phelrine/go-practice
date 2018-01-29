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

	f := struct {
		f1 int
		f2 string
	}{
		f1: 1,
		f2: "struct",
	}
	fmt.Printf("struct f = { f1 = %d, f2 = %s }\n", f.f1, f.f2)

	type myType struct {
		num  int
		name string
	}
	g := myType{
		10,
		"name1",
	}
	fmt.Printf("myType g = { num = %d, name = %s }\n", g.num, g.name)
}
