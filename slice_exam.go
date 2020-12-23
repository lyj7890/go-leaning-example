package main

import "fmt"

func main() {
	var s1 []int = make([]int, 3, 5)
	s1 = []int{1, 2, 3}
	//var s2 []int
	s2 := s1[:]
	s1 = append(s1, 4)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	//assignment(s1, s2)
}

func appSli(s []int, v int) {
	fmt.Println("before append:", s, "length:", len(s), "capacity:", cap(s))
	_ = append(s, v)
	fmt.Println("after append:", s, "length:", len(s), "capacity:", cap(s))
}

func assignment(s1, s2 []int) {
	fmt.Println("before assign, s1:", s1, "s2:", s2)
	s2 = s1
	s1[0] = 100
	fmt.Println("after assign, s1:", s1, "s2:", s2)
}
