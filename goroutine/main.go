package main

import (
	"fmt"
	"time"
)

func main() {
	finished := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(i2 int) {
			fmt.Printf("%d start\n", i2)
			if i2 == 9 {
				time.Sleep(time.Second * 2)
			} else {
				time.Sleep(time.Second)
			}
			fmt.Printf("%d done\n", i2)
			if i2 == 9 {
				finished <- true
			}
		}(i)
	}
	<-finished
	close(finished)
}
