package main

import "fmt"

func main() {
	dict := map[int]string{
		90: "优秀",
		80: "良",
		60: "一般",
		30: "差",
	}
	if res,ok := dict[91]; ok {
		fmt.Println("exist", res)
	} else {
		fmt.Println("not exist", res)
	}
}