package main

import (
	"encoding/json"
	"fmt"
)

type response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	type User struct {
		Name string `json:"name"`
	}

	type Userinfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var userResponse response[User]
	json.Unmarshal([]byte(`{"code":1,"msg":"瑶瑶真棒","data":{"name":"瑶瑶"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)

	var userinfoResponse response[Userinfo]
	json.Unmarshal([]byte(`{"code":1,"msg":"瑶瑶真棒","data":{"name":"杰杰","age":24}}`), &userinfoResponse)
	fmt.Println(userinfoResponse.Data.Name, userinfoResponse.Data.Age)
}
