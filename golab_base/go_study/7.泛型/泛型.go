package main

import (
	"encoding/json"
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func plus[T Number](n1, n2 T) T {
	return n1 + n2
}

func myPrint[T int, K string | int](n1 T, n2 K) {
	return
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func main() {
	//plus(1, 2)
	//var u1, u2 = uint(2), uint(4)
	//x := plus(u1, u2)
	//fmt.Println(x)
	type User struct {
		Name string `json:"name"`
	}

	type Userinfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	userres := Response{
		Code: 1,
		Msg:  "瑶瑶真棒",
		Data: User{
			Name: "瑶瑶",
		},
	}
	byteData1, _ := json.Marshal(userres)
	fmt.Println(string(byteData1))

	userinfores := Response{
		Code: 1,
		Msg:  "瑶瑶真棒",
		Data: Userinfo{
			Name: "瑶瑶",
			Age:  12,
		},
	}
	byteData2, _ := json.Marshal(userinfores)
	fmt.Println(string(byteData2))

}
