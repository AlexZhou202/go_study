package main

import "fmt"

func main() {
	var age = 30
	var name = "zjj"
	fmt.Println("后年:",age)
	fmt.Println("打印第一年")
	fmt.Println("打印第二年")
	fmt.Print("打印第一年")
	fmt.Print("打印第二年")
	fmt.Printf("\n%T,%s,%T,%d\n",name,name,age,age)  //占位符：%T 打印变量类型,%s 打印变量为字符,%d 打印类型为整数,\n 回车换行
}
