package leaning

import (
	"fmt"
	"sync"
)

func main(){
	p := &Person{}
	fmt.Printf("%s%+v\n","before update:",p)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i < 10;i++ {
		go p.update()
		wg.Done()
	}
	wg.Wait()
	fmt.Printf("%s%+v\n","after update:",p)
}

type Person struct {
	sync.Mutex
	Name string
	age int
}

//change person
func (p *Person)update(){
	p.Lock()
	p.Name = "golang"
	p.age += 1
	p.Unlock()
}