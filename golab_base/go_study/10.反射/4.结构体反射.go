package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int
	IsMan bool
}

// 语法：
// reflect.ValueOf(obj)：获取对象的反射值，可以用来访问字段的值。
// reflect.TypeOf(obj)：获取对象的类型信息，可以用来访问字段的元数据（如标签）。
// v.NumField()：返回结构体字段的数量，用于遍历所有字段。
// t.Field(i)：获取第 i 个字段的类型信息（包括标签）。
// tf.Tag.Get("json")：获取字段的 json 标签内容。
// v.Field(i)：获取第 i 个字段的实际值
func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = tf.Name
		}
		//fmt.Printf("%s,%#v\n", tf.Name, jsonTag)
		fmt.Println(jsonTag, v.Field(i))
	}
}
func main() {
	s := Student{Name: "瑶瑶", Age: 18, IsMan: false}
	ParseJson(s)
}
