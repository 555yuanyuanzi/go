package main

import "fmt"

// 函数参数定义
func say() {
	fmt.Println("hello")
}
func param1(id, name string) {
	fmt.Println(id, name)
}
func add1(list ...int) { //...意思是展开切片
	var sum int
	for _, value := range list {
		sum += value
	}
	fmt.Println(sum)
}

func add2(list []int) { //...意思是展开切片
	var sum int
	for _, value := range list {
		sum += value
	}
	fmt.Println(sum)
}

// 无返回值
func r1() {
	return
}

// 单返回值
func r2() bool {
	var ok bool
	return ok
}

// 多返回值
func r3() (string, bool) {
	return "", true
}

// 命名返回值，相当于先定义再赋值
func r4() (val string, ok bool) {
	if 5 > 2 {
		val = "12"
		//ok如果不赋值，会默认为false
		//ok = true
		return
	}
	return
}

func main() {

	//say()
	//param1("1001", "小杰")
	/*
		list := []int{1, 2}
		add1(4, 5)
		add1(list...)
		add2(list)
		fmt.Println(r2())
		fmt.Println(r4())
	*/
	//创建匿名函数，需要在函数里面定义函数
	var setname = func(name string) {
		fmt.Println(name)
		return
	}
	setname("张三")

}
