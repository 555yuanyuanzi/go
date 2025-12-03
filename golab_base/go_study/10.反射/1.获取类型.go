package main

import (
	"fmt"
	"reflect"
)

func getType(obj any) {
	//TypeOf获取类型，Kind获取实际类型的值
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("int")
	case reflect.Struct:
		fmt.Println("struct")
	}
}

func main() {
	getType(123)
	getType("23")
	getType(struct {
		Name string
	}{})
}
