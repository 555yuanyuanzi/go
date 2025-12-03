package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Println("测试前")
}
func teardown() {
	fmt.Println("测试后")
}
func TestADD2(t *testing.T) {
	fmt.Println("测试中")
	t.Error("哈哈哈哈")
}

// 测试主入口
func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
