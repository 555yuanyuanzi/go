package main

import (
	"fmt"
	"net"
	"time"
)

//服务器	监听端口 → 等连接 → 发送 hello world → 关闭连接
//客户端	连服务器 → 接收数据 → 打印 → 看到 EOF 后退出

// 写一个最简单的 TCP 服务器
// 服务器监听在 127.0.0.1:81，等待客户端连接。
// 每有一个客户端连上来，就给它回一句 "hello world"，然后断开。
func main() {
	// 创建tcp的监听地址
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:81")
	// tcp监听
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	//提示tcp服务器绑定的地址是 127.0.0.1:81
	fmt.Println("tcp server bind addr :", addr.String())
	for {
		// 等待客户端连接，accept是阻塞的
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		//打印客户端地址
		fmt.Println(conn.RemoteAddr())
		//给客户端发消息，往 TCP 连接里写 14 个字节
		conn.Write([]byte("hello world"))
		time.Sleep(2 * time.Second)
		conn.Close() //服务端关闭连接
	}
}
