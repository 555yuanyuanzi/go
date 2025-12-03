package main

import "fmt"

func Copy(name string) {
	fmt.Printf("infunc %p\n", &name)
	name = "小瑶瑶"
}
func set(name *string) {
	fmt.Printf("infunc %p\n", name) //此时name是存的地址
	*name = "瑶瑶不知道"
	//解引用，通过内存去找值
	//&是取地址，*是解引用，去这个地址指向的值

}
func main() {
	name := " 瑶瑶"
	// 函数参数传值，拷贝一份新的变量传入函数
	fmt.Printf("outfunc %p\n", &name)
	//Copy(name)
	set(&name)
	fmt.Println(name)

}
