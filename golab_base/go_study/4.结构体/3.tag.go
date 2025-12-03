package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Psssword string `json:"-"`
	//"-"表示转json的时候忽略该字段,"omitempty"表示忽略空值
}

func main() {
	user := User{Name: "瑶瑶", Age: 18, Psssword: "142106"}
	//json.Marshal 返回两个值：第一个是 []byte（UTF-8 编码的 JSON 数据），第二个是 error
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}
