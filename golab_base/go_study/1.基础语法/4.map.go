package main

import "fmt"

func main() {
	var user_map map[int]string = map[int]string{
		1: "yaoyao",
		2: "jiejie",
		3: "yuanyunan",
		4: "",
	}
	fmt.Println(user_map[3])
	fmt.Printf("%#v\n", user_map[5])
	value := user_map[2]
	fmt.Println(value)
	// 如何判断有没有取到值
	value, ok := user_map[3]
	fmt.Println(value, ok)
	// map取值，如果只有一个参数杰，那这个参数就是值，如果没有，这个参数类型就是零值
	//如果两个参数接，那第二个参数就是布尔值，表示是否有这个元素

	// 添加和修改
	user_map[1] = "yaoyao走了"
	fmt.Println(user_map)
	delete(user_map, 3)
	fmt.Println(user_map)

	//map初始化
	//var amap = make(map[string]string)
	var amap = map[string]string{}
	amap["name"] = "yaoyao"
	fmt.Println(amap)

}
