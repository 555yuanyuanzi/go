package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("event执行开始")
	time.Sleep(2 * time.Second)
	fmt.Println("event执行结束")
	close(done)
}

func main() {
	go event()
	select {
	case <-done:
		fmt.Println("协程执行完毕")
	case <-time.After(1 * time.Second):
		fmt.Println("超时")
		return
		// time.After(1s) 会在 1 秒后 自动写入一个时间值
		//这是 Go 内置的一种“定时器通道”
	}

}
