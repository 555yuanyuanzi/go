package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{}
	var wait sync.WaitGroup
	// 三个协程实现写maps操作
	wait.Add(3)
	go func() {
		defer wait.Done()
		maps.Store(1, "瑶瑶")
	}()
	go func() {
		defer wait.Done()
		maps.Store(2, "杰杰")
	}()
	go func() {
		defer wait.Done()
		maps.Store(3, "源源子")
	}()
	wait.Wait()

	found := false
	// 遍历maps
	maps.Range(func(key any, value any) bool {
		//fmt.Println(key, value)
		//if 初始化语句; 条件判断 {
		//	// 代码
		//}
		if value, ok := value.(string); ok && value == "杰杰" {
			fmt.Println("找到杰杰了！！")
			found = true
			return false
		}
		return true // 继续遍历下一个元素的意思
	})

	fmt.Println("找到人了", found, "结束")
}
