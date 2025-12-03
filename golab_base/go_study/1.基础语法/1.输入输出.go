package main

import (
	"fmt"
	"sort"
)

func main() {
	var ints = []int{4,8,3,2}
	sort.Ints(ints)
	fmt.Println("升序：",ints)
	sort.Sort(sort.Reverse(sort.IntSlice((ints))))
	fmt.Println("降序:",ints)
	//测试输出格式
	fmt.Printf("%.3f\n",2.12231)
	fmt.Printf("%T %T\n","你好",12)//%T打印类型
	fmt.Printf("%#v\n","")//%#v打印值的详细信息（打印了空字符串还是没有打印）
	fmt.Printf("%v\n","")
	var f = fmt.Sprintf("%.2f",2.12231)//Sprintf返回一个格式化的字符串
	fmt.Println(f)

	//输入，用的不多
	fmt.Println("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Println("请输入你的年龄")
	var age int
	n,err := fmt.Scan(&age)
	fmt.Println(n,err,age)
}