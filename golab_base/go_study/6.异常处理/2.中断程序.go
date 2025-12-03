package main

import (
	"fmt"
	"os"
)

// 遇到错误直接停止程序，这种一般是用于初始化，一旦初始化出现错误，程序继续走下去意义也不大了，还不如直接中断掉
func init() {
	_, err := os.ReadFile("1234")
	if err != nil {
		//log.Fatalln("错误了")
		panic("错误了")
	}
}

func main() {
	fmt.Println("main")
}
