package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 一次性读取
	/*
		byteData, err := os.ReadFile("go_study/8.文件/hello.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(byteData))
	*/

	// 分片读
	/*
		file, err := os.Open("go_study/8.文件/hello.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		for {
			var byteData = make([]byte, 24)
			n, err := file.Read(byteData)
			if err == io.EOF {
				break
			}
			fmt.Println(string(byteData[:n]))
		}
	*/

	// 按行读
	file, err := os.Open("go_study/8.文件/hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//fmt.Println("------按行读：")
	// buf := bufio.NewReader(file) // 创建缓冲对象
	//for {
	//	line, _, err := buf.ReadLine()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(line), err)
	//}

	buf2 := bufio.NewScanner(file)
	buf2.Split(bufio.ScanWords)
	var index int
	for buf2.Scan() {
		index++
		fmt.Println(index, buf2.Text())
	}

}
