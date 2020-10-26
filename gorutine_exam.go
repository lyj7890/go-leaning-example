package main

import (
	"fmt"
)

func main() {
	fmt.Println("main goroutine run")
	ch := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			ch <- i
		}
		close(ch)
	}()

	for{
		i,ok := <- ch
		if !ok {
			fmt.Println("channel close")
			break
		}
		fmt.Println(i,ok)
	}


	//for i := range ch {
	//	fmt.Println(i)
	//}

}
