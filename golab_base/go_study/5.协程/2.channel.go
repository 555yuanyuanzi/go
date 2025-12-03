package main

import (
	"fmt"
	"sync"
	"time"
)

// 在多个 goroutine 里同时读写 sync.Map 是安全的，不会崩。
var moneyChannel = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("---%s 开始购物 \n", name)
	// 购物时间
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束 \n", name)

	// 发送（写入
	moneyChannel <- money // 把数据写入信道
	wait.Done()
}

// 引出问题，在协程里面产生的数据怎么传递给主线程或者其他协程函数呢？---channel
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)
	go pay("妖妖", 2, &wait)
	go pay("源源子", 3, &wait)
	go pay("杰杰", 5, &wait)

	go func() {
		// 确保在所有写入完成后再关闭通道
		defer close(moneyChannel) //关闭信道
		wait.Wait()
		//close(moneyChannel)
	}()

	var moneyList []int

	for money := range moneyChannel {
		//fmt.Println(money)
		moneyList = append(moneyList, money)
	}
	/*
		for {
			// 接收（读取）
			money, ok := <-moneyChannel
			fmt.Println(money, ok)
			if !ok {
				break
			}
		}
	*/
	fmt.Println("所有人购物完成", time.Since(startTime))
	fmt.Println("花钱清单：", moneyList)
}
