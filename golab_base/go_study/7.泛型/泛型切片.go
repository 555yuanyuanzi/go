package main

import "fmt"

func main() {
	/*
		// 泛型切片
		type MySlice[T int | string] []T
		var mySlice = MySlice[int]{1, 2, 3}
		fmt.Println(mySlice[0] + 4)
	*/

	// 泛型map,map的key只能是基本数据类型
	type MyMap[T string, K any] map[T]K
	var myMap = MyMap[string, int]{
		"name": 12,
	}
	fmt.Println(myMap)

}
