package main

import "fmt"

func main() {
	dict := map[int]string{
		90: "优秀",
		80: "良",
		60: "一般",
		30: "差",
	}
	fmt.Println("dict:", dict, "length:", len(dict))
	fmt.Println("read dict[90]:", dict[90])
	fmt.Println("dict[10]", dict[10])
	dict[10] = "糟糕"
	fmt.Println("dict[10]", dict[10])
	dict[90] = "很好"
	fmt.Println("dict[90]",dict[90])
	fmt.Println("dict:", dict, "length:", len(dict))
}
