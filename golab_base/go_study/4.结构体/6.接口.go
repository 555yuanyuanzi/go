//非常重要！！

package main

import "fmt"

// 定义一个接口,实现唱歌，getname的方法
type SingInterface interface {
	Sing()
	GetName() string
	Xingbie() string
}

type chicken struct {
	Name string
}
type cat struct {
	Name string
}

func (c chicken) Sing() {
	fmt.Println(c.Name + "在唱歌")
}
func (c chicken) GetName() string {
	return c.Name
}
func (c chicken) Xingbie() string {
	return "男生"
}

func (c cat) Sing() {
	fmt.Println(c.Name + "在唱歌")
}
func (c cat) GetName() string {
	return c.Name
}

func (c cat) Xingbie() string {
	return "女生"
}

// 注意：结构体类型必须满足接口的所有方法，才能实现该接口，作为参数传递
func Sing(c SingInterface) {
	// 单一类型断言
	//ch, ok := c.(chicken)
	//fmt.Println(ch, ok)

	// 多种类型断言
	switch server := c.(type) {
	case chicken:
		fmt.Printf("%#v\n", server)
	case cat:
		fmt.Println(server)
	default:
		fmt.Println("其他类型")
	}
	//c.Sing()
	//fmt.Println("--", c.Xingbie(), c.GetName())
}

// 空接口可以接受任何类型，因为任何类型都实现了空接口
type EmptyInterface interface {
}

// 封装打印函数
// 写法1：
func Print(val EmptyInterface) {
	fmt.Println(val)
}

// 写法2:
func Print2(val interface{}) {
	fmt.Println(val)
}

// 写法3：
func Print3(val any) {
	fmt.Println(val)
}

func main() {
	ch := chicken{Name: "瑶瑶"}
	ca := cat{Name: "阿杰"}
	Sing(ch)
	Sing(ca)
	Print(ch)
	Print2(ch)
	Print(ca)
	Print3(ca)
}
