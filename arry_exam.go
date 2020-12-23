package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	//var s1 []int = []int{1,2,3,4,5}
	var arr2 = [0]int{}
	fmt.Println(arr2)

	arr1[0] = 100
	fmt.Println(arr1)
	for i, v := range arr1 {
		fmt.Printf("index %d is %d\n",i, v)
	}
	for m := range arr1 {
		fmt.Printf("index %d is %d\n", m, arr1[m])
	}
}

func squares(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
}
