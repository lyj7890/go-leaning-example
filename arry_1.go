package main

import "fmt"

func main() {
	arr := [...]string{1:"周一",2:"周二",3:"周三",4:"周四",5:"周五"} //初始化数组时可以指定索引，未指定的索引值为零值
	fmt.Println(arr[0])  //索引0处未指定，为零值。
	s := [...]int{1,2,3,4,5}
	fmt.Println("before rev:",s)
	rev(s[:])
	fmt.Println("after rev:",s)
}

//使用slice翻转数组
func rev(s []int){
	for i,j := 0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j]=s[j],s[i]
	}
}