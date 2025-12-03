package main

import "fmt"

func main() {
	// 传统for循环
	/*
		var sum int
		for i := 0; i < 10; i += 2 {
			sum += i
		}
		fmt.Println(sum)
	*/

	//死循环
	/*
		for {
			fmt.Println(time.Now())
			time.Sleep(2 * time.Second)
		}
	*/

	//while模式的循环(先判断条件，再执行循环体)
	/*
		var sum int
		var i int = 1
		for i <= 100 {
			sum += i
			i++
		}
		fmt.Println(sum)
	*/

	//do while模式
	/*
		var sum int
		var i int = 1
		for {
			sum += i
			i++
			if i > 100 {
				break
			}
		}
		fmt.Println(sum)
	*/

	//循环切片(列表)
	/*
		var list = []string{"瑶瑶", "杰杰"}
		for i := 0; i < len(list); i++ {
			fmt.Println(i, list[i])
		}
		for index, item := range list {
			fmt.Println(index, item)
		}

		//循环map
		var userMap = map[string]string{
			"name": "瑶瑶",
			"age":  "18",
		}
		for key, value := range userMap {
			fmt.Println(key, value)
		}
	*/

	// 打印99乘法表
	for i := 1; i <= 9; i++ {
		for j := i; j <= 9; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println() //打印换行
	}
	//continue:跳过本次循环，进行下一次循环
	//break:结束整个循环
}
