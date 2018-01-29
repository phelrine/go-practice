package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 算術演算
	fmt.Printf("2 + 3 = %d\n", 2+3)
	fmt.Printf("4 %% 2 = %d\n", 4%2)

	// 文字列演算
	str1 := "abc" + "def"
	fmt.Printf("str1 = %s\n", str1)
	str1 += "ghi"
	fmt.Printf("str1 = %s\n", str1)
	ch := str1[2]
	fmt.Printf("str1[2] = %c\n", ch)
	fmt.Printf("len(str1) = %d\n", len(str1))

	// 日本語
	str2 := "テスト"
	fmt.Printf("len(\"テスト\") = %d\n", len(str2))
	fmt.Printf("len(\"テスト\") = %d\n", utf8.RuneCountInString(str2))
}
