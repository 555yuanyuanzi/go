package main

import "fmt"

//	type 结构体名称 struct{
//		名称 类型//成员或属性
//	}
type class struct {
	className string
}
type student struct {
	class //继承class结构体
	Name  string
	Age   int
	Job   string
}

func (s student) Study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}
func (s student) Info() {
	fmt.Printf("名字: %s 班级: %s\n", s.Name, s.className)
}

// 指针可以修改结构体里面的内容
func (s *student) SetName(name string) {
	s.Name = name
}
func main() {
	c1 := class{className: "go语言班"}
	fmt.Println("学生1：")
	s1 := student{class: c1, Name: "yaoyao", Age: 18, Job: "算法工程师"}
	s1.Study()
	s1.SetName("阿杰")
	s1.Info()
	fmt.Println(s1)
	fmt.Println("-------------------")
	fmt.Println("学生2：")
	s2 := student{class: c1, Name: "jiejie", Age: 25, Job: "开发工程师"}
	s2.Study()
	s2.Info()

}
