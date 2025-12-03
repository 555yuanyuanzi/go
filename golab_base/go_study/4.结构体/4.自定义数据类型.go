package main

import "fmt"

//1001 1002 1003 1004

type Code int

// 为Code类型定义一个方法GetMsg
func (code Code) GetMsg() string {
	switch code {
	case SuccessCode:
		return "成功"
	case ServiceErrCode:
		return "服务错误"
	case NetworkErrCode:
		return "网络错误"
	}
	return ""
}
func (code Code) OK() (c Code, msg string) {
	return code, code.GetMsg()
}

const (
	SuccessCode    Code = 0    //
	ServiceErrCode Code = 1001 //服务错误
	NetworkErrCode Code = 1002 // 网络错误
)

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.OK()
	}
	if name == "2" {
		return NetworkErrCode.OK()
	}
	return SuccessCode.OK()
}
func main() {
	code, msg := webServer("9")
	fmt.Println(code, msg)
}
