package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	//中断式，卫语句（最好的）
	/*
		if age <= 18 {
			fmt.Println("未成年")
			return
		}
		if age <= 35 {
			fmt.Println("青年")
			return
		}
		fmt.Println("中年")
	*/
	//嵌套式
	if age <= 18 {
		fmt.Println("未成年")
	} else {
		if age <= 35 {
			fmt.Println("青年")
		} else {
			fmt.Println("中年")
		}
	}
	//逻辑短路
	//&&第一个条件如果是false，后面的条件就不会走了
	//||第一个条件如果是true，后面的条件就不会走了
}
