package main

import (
	"fmt"
	"sync"
	"time"
)

var MoneyChannel = make(chan int)     // 花了多少钱
var NameChannel = make(chan string)   // 谁花的钱
var doneChannel = make(chan struct{}) // 结束信号，不需要传数据，struct{} 是 Go 里“最小、零开销”的类型

func buy(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("---%s 开始购物 \n", name)
	// 购物时间
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束 \n", name)

	// 发送（写入
	MoneyChannel <- money // 把数据写入信道
	NameChannel <- name
	wait.Done()
}

// 引出问题，在协程里面产生的数据怎么传递给主线程或者其他协程函数呢？---channel
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)
	go buy("妖妖", 2, &wait)
	go buy("源源子", 3, &wait)
	go buy("杰杰", 5, &wait)

	go func() {
		// defer后进先出
		// 确保在所有写入完成后再关闭通道
		defer close(MoneyChannel)
		defer close(NameChannel) //关闭信道
		defer close(doneChannel)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string

	var event = func() {
		for {
			select {
			// 从 MoneyChannel 读取一个值，如果信道里有值就立即拿到；如果没有值就阻塞等待。
			// 拿到的值放到变量 money 里，然后执行这个 case 的代码。
			//阻塞等待，直到 channel 有一个值可以读出来,读到后赋值给 x
			case money := <-MoneyChannel:
				moneyList = append(moneyList, money)
			case name := <-NameChannel:
				nameList = append(nameList, name)
			// 只有在close(doneChannel)执行后，doneChannel 才会变成“已关闭状态”，是“下一次 select 会立刻执行”
			case <-doneChannel:
				return
			}
		}
	}
	event()
	fmt.Println("所有人购物完成", time.Since(startTime))
	fmt.Println("花钱清单：", moneyList)
	fmt.Println("姓名清单：", nameList)
}
