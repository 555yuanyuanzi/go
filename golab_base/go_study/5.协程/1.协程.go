package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局变量写法，也可以作为指针传参
//var wait sync.WaitGroup

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("---%s 开始购物 \n", name)
	// 购物时间
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束 \n", name)
	wait.Done()
}

// 协程
// 引出问题，在协程里面产生的数据怎么传递给主线程或者其他协程函数呢？---channel
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	// 购物接力模式
	//shopping("妖妖")
	//shopping("源源子")
	//shopping("杰杰")

	// 协程函数数目
	wait.Add(3)
	// 主线程结束，协程函数跟着结束
	go shopping("妖妖", &wait)
	go shopping("源源子", &wait)
	go shopping("杰杰", &wait)

	wait.Wait()
	// 把等待时间写死：
	//time.Sleep(1 * time.Second)
	fmt.Println("所有人购物完成", time.Since(startTime))
}
