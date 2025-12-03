package main

import "fmt"

func main() {
	fmt.Println("请输入你要执行的操作：")
	fmt.Println("1.登录 2.注册 3.个人中心")
	var index int
	fmt.Scan(&index)

	var funcMap = map[int]func(){
		1: login,
		2: register,
		3: userCenter,
	}
	fun, ok := funcMap[index]
	if ok {
		fun()
	}
}
func login() {
	fmt.Println("登录")
}

func register() {
	fmt.Println("注册")
}

func userCenter() {
	fmt.Println("个人中心")

}
