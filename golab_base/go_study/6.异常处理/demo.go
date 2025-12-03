package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	fmt.Println("init() 已执行")
}
func read1() {
	list := []int{1, 2}
	fmt.Println(list[4]) // 越界：panic
}

func main() {
	fmt.Println("main 开始（最先的打印）")
	// 在 main 里用 defer + recover 兜底
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到 panic：", err)
			fmt.Println("堆栈信息：")
			fmt.Println(string(debug.Stack()))
		}
	}()
	fmt.Println("开始执行read:")

	read1() // 这里会 panic，中断main的正常流程，但被上面的 defer/recover 抓住
	// 后面的代码不会执行
	fmt.Println("正常逻辑")
}
