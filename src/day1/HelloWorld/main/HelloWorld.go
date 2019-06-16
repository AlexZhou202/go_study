package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HelloWorld")
	time.Sleep(5 * time.Second) //加了等待5秒防止windows编译成exe双击一闪而过的问题
}
