package main

import "fmt"

//类型别名，和自定义类型很想，但是有一些地方和自定义类型有很大差异
//1.不能绑定方法
//2.打印类型还是原始类型
//3.和原始类型比较，类型别名不用转换(类型别名就是int ，但是MyCode不是int)

type MyCode int        //自定义类型
type MyAliasCode = int //类型别名

func (m MyCode) Code() {

}

const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode) //1,main.MyCode
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode)
	age := 1
	if int(myCode) == age {
		fmt.Println("类型别名可以和原始类型比较")
	}
}
