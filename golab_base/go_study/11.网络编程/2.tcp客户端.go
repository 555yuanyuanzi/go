package main

import (
	"fmt"
	"io"
	"net"
)

// 这个客户端程序会连接你的 TCP 服务器（127.0.0.1:81），
// 然后不断读取服务器发来的数据，打印出来，
// 直到服务器关闭连接。
func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:81")
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建一个无限循环，持续读取服务器的数据
	for {
		//创建一个 1024 字节的缓冲区
		var byteData = make([]byte, 1024)
		//从 conn 中读数据
		//Read() 会 阻塞等待，直到服务器发送数据
		n, err := conn.Read(byteData)
		//判断服务器是否关闭连接,err == io.EOF 是判断“是否读到文件/连接的结束
		if err == io.EOF {
			break
		}
		fmt.Println(string(byteData[:n]))
	}
}
