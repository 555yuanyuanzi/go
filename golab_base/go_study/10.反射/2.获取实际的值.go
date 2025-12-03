package main

import (
	"fmt"
	"reflect"
)

func getValue(obj any) {
	//TypeOf获取类型，Kind获取实际类型的值
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("string", v.String())
	case reflect.Int:
		fmt.Println("int", v.Int())
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	getValue(123)
	getValue("23")
	getValue(struct{ Name string }{Name: "瑶瑶"})
}
