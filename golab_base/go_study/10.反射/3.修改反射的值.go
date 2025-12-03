package main

import (
	"fmt"
	"reflect"
)

// 修改obj的值为value
func setValue(obj any, value any) {
	//TypeOf获取类型，Kind获取实际类型的值
	//reflect.TypeOf(x) 得到的是 reflect.Type 类型，表示变量的类型信息（比如 int、string、struct 等）。
	//reflect.ValueOf(x) 得到的是 reflect.Value 类型，表示变量的值信息，可以用来读写、修改实际的值。
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)
	// v1.Kind()拿到的是指针类型，需要用Elem()解引用
	fmt.Println(v1.Elem().Kind())
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() {
	case reflect.String:
		// 类型断言只能用于接口类型
		// reflect.Value 类型没有 String() 方法来直接获取字符串值
		v1.Elem().SetString(value.(string))
	case reflect.Int:
		//v1.Elem().SetInt(v2.Int())
		v1.Elem().SetInt(int64(value.(int)))
	}
}

func main() {
	var name = "瑶瑶"
	var age = 23
	setValue(&name, "源源")
	setValue(&age, 18)
	fmt.Println(name, age)

}
