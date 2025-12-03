package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	list := []int{1, 2}
	fmt.Println(list[3])
}
func main1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			// 找错误信息的位置，打印堆栈信息
			fmt.Println(string(debug.Stack()))
		}
	}()
	read() // panic中断了main1函数的正常流程
	fmt.Println("main1没有被中断")
}

func main() {
	main1() // 3. 调用 main1。因为 panic 被 main1 内部的 recover 处理了，
	//    所以 main1 能够正常返回，没有把 panic 传递给 main。
	// 4. 由于 main1 正常返回了，main 函数的流程继续向下执行
	fmt.Println("正常逻辑")
}
