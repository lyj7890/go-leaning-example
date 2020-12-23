<<<<<<< HEAD
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
=======
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
>>>>>>> 36dfdc5f5461d1e2ce9537785c0720ee4119bd4e
