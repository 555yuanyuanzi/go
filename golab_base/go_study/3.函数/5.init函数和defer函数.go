package main

import "fmt"

//init（）函数有以下特性：
//1.一个包可以有多个init函数，init函数可以出现在任意文件中
//2.init函数不能被其他函数调用，init函数没有参数也没有返回值
//3.init函数在main函数执行前被调用，且优先于main函数执行
//4.不能作为参数传入

// defer()函数：用于延迟执行某个操作，通常用于资源释放、文件关闭、解锁等场景
// defer语句会将其后面的函数调用延迟到包含它的函数执行结束时才执行
// defer语句遵循先进后出（LIFO）的原则，即最后一个defer语句会最先被执行
// defer语句中的变量在defer声明前就被定义好了
var db int

func init() {
	db = 10
	fmt.Println("连接数据库")
	fmt.Println("init-1")
}
func init() {
	fmt.Println("init-2")
}

func main() {
	//fmt.Println("main", db)
	age := 18
	defer func() {
		fmt.Println("defer-1")
		fmt.Println(age)
	}()
	defer fmt.Println("defer-2")
}
