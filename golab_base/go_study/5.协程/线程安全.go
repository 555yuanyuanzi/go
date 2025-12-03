package main

import (
	"fmt"
	"sync"
)

var sum int
var wait sync.WaitGroup
var lock sync.Mutex

func add() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		sum++
	}
	wait.Done()
	lock.Unlock()
}
func sub() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		sum--
	}
	wait.Done()
	lock.Unlock()
}

func main() {
	wait.Add(2)
	go add()
	go sub()
	wait.Wait()
	fmt.Println(sum)
}
