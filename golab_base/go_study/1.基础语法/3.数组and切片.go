package main

import (
	"fmt"
	"sort"
)

func main(){
	// 定义数组（长度限制）
	var nameList [3]string = [3]string{"张三","李四","王五"}
	fmt.Println(nameList)
	fmt.Println((nameList[0]))
	fmt.Println(len(nameList))
	// 错误：fmt.Println(nameList[-1])go不支持负向索引

	// 定义切片（没有长度限制，可以动态变化）
	var list []string
	//list := make([]string, 0)
	fmt.Println(list==nil) //如果只定义不赋值，那么实际值就是nil
	// make函数可以创建指定长度和容量的切片
	list = append(list,"小明")
	list = append(list,"瑶瑶")
	fmt.Println(list)
	list[1] = "小杰"
	fmt.Println(list)

	// 从数组得到切片
	array := [4]int{1,2,3,4}
	slices := array[0:2]//左闭右开
	fmt.Println(slices)	

	// 切片排序
	var ints = []int{4,8,2,3}
	sort.Ints(ints)
	fmt.Println(ints)//升序排序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)//降序排序

}