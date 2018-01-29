package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	arr := []int{10, 3, 8, 7, 5, 2, 1, 6, 4}
	var g errgroup.Group
	for _, i := range arr {
		j := i
		g.Go(func() error {
			time.Sleep(time.Second * time.Duration(j))
			fmt.Printf("%d\n", j)
			return nil
		})
	}
	_ = g.Wait()
}
