package main

import (
	"fmt"
)

func main() {
	// 構造体
	type myStruct struct {
		field1 string
		field2 int
	}
	s1 := myStruct{
		field1: "f1",
		field2: 2,
	}
	fmt.Printf("s1.field1 = %s\n", s1.field1)
	fmt.Printf("s1.field2 = %d\n", s1.field2)

	// スライス
	var slice1 []int
	// 10個の要素を持つスライスを作成
	slice1 = make([]int, 10)
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i
	}
	// 要素列挙
	for i, e := range slice1 {
		fmt.Printf("slice1[%d] = %d\n", i, e)
	}
	slice1 = append(slice1, 100)
	fmt.Printf("len(slice1) = %d\n", len(slice1))
	fmt.Printf("%d\n", slice1[10])

	slice2 := []int{}
	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i)
	}

	// スライスの連結
	slice3 := append(slice1, slice2...)
	for _, e := range slice3 {
		fmt.Printf("%d\n", e)
	}

	map1 := map[string]int{}
	map1["abc"] = 1
	map1["123"] = 2

	// mapのサイズ取得
	fmt.Printf("len(map1) = %d\n", len(map1))

	// 要素アクセス
	fmt.Printf("map[\"abc\"] = %d\n", map1["abc"])

	// キーチェック
	if e, ok := map1["xxx"]; ok {
		fmt.Printf("key is exists %d\n", e)
	} else {
		fmt.Println("key is not exists")
	}

	// 要素列挙
	for k, v := range map1 {
		fmt.Printf("map1[%s] = %d\n", k, v)
	}

	// キーのみ
	for k := range map1 {
		fmt.Printf("key is %s\n", k)
	}
}
