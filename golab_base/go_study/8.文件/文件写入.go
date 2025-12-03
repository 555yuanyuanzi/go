package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		file, err := os.OpenFile("go_study/8.文件/w.txt", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(byteData))
	*/

	// 全部写入
	//err := os.WriteFile("w1.txt", []byte("我喜欢你"), 0666)
	//fmt.Println(err)

	//文件拷贝
	/*
		rFile, err := os.Open("C:\\Users\\86155\\Pictures\\锁屏\\QQ图片20251030212952.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rFile.Close()
		wFile, err := os.OpenFile("zeng.jpg", os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer wFile.Close()

		io.Copy(wFile, rFile)
	*/
	dir, err := os.ReadDir("go_study/8.文件")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
}
