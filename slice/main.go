package main

import (
	"fmt"
)

func main() {
	var s1 []int
	s1 = []int{1, 2, 3, 4, 5}

	fmt.Printf("s1[0] = %d\n", s1[0])
	fmt.Printf("s1[1] = %d\n", s1[1])
	fmt.Printf("s1[2] = %d\n", s1[2])
	fmt.Printf("s1[3] = %d\n", s1[3])
	fmt.Printf("s1[4] = %d\n", s1[4])
	fmt.Printf("len(s1) = %d\n", len(s1))

	s2 := make([]int, 3)
	s2[0] = 10
	s2[1] = 20
	s2[2] = 30
	fmt.Printf("s2[0] = %d\n", s2[0])
	fmt.Printf("s2[1] = %d\n", s2[1])
	fmt.Printf("s2[2] = %d\n", s2[2])
	fmt.Printf("len(s2) = %d\n", len(s2))

	var s3 []int
	fmt.Printf("len(s3) = %d\n", len(s3))
	s3 = append(s3, -1)
	fmt.Printf("len(s3) = %d\n", len(s3))
	s3 = append(s3, -2)
	fmt.Printf("len(s3) = %d\n", len(s3))
	s3 = append(s3, -3)
	fmt.Printf("len(s3) = %d\n", len(s3))
	fmt.Printf("s3[0] = %d\n", s3[0])
	fmt.Printf("s3[1] = %d\n", s3[1])
	fmt.Printf("s3[2] = %d\n", s3[2])
}
