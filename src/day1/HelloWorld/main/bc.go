package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue //continue跳出本次循环
		}
		fmt.Println(i)
	}

	for i := 0; i <= 5; i++ {
		if i == 3 {
			break //break跳出循环（结束循环）
		}
		fmt.Println(i)
	}
}
