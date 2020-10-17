package main

import "fmt"

type AA struct {
	Name string
}
type BB struct {
	Name string
}

//接收者为AA,值传递
func (a AA) Print(name string) {
	//fmt.Println(a.Name)
	a.Name = "ppjj"
	fmt.Println(a.Name + name)

}
//接收者为BB,引用传递
func (a *BB) Print(name string) {
	a.Name = "ppjj"
	fmt.Println(a.Name + name)

}
func main() {
	//值传递，不能修改字段
	a := AA{Name: "pj"}
	a.Print("test")
	fmt.Println(a.Name)
	//引用传递，可以修改字段
	b := BB{Name: "pj"}
	b.Print("test")
	fmt.Println(b.Name)
}
