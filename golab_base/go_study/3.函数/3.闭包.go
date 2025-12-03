package main

import (
	"fmt"
	"time"
)

// 闭包，设计一个函数，先传一个参数表示延时，后面再次传参数就是将参数求和
func awaitADD(awaitSecond int) func(...int) int {
	//延时
	time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberlist ...int) (sum int) {
		//用到了waitSecond变量，表示闭包
		//time.Sleep(time.Duration(awaitSecond) * time.Second)
		for _, i2 := range numberlist {
			sum += i2
		}
		return sum
	}
}

func main() {
	add := awaitADD(2)
	t1 := time.Now()
	sum := add(1, 2, 3)
	subTime := time.Since(t1)
	fmt.Println(sum, subTime)

}
