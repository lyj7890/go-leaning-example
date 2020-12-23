package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main goroutine run")
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println(i, "goroutine running")
			for {}
		}()
	}
	for {
		fmt.Println("in main goroutine")
		time.Sleep(time.Second)
	}

}
