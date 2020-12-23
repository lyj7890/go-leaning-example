package main

import "fmt"

func main(){
	arr := []int{1,2}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _,v := range arr {
		count[v]++
	}
	times := []int{}
	for _,v := range count {
		times = append(times,v)

	}
	fmt.Println(times)
	return len(times) == len(count)
}