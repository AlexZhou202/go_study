package main

import "fmt"

func main() {
	A := 2
	b := A
	b = 3
	fmt.Println(A, b)
	//指针，不是单纯的赋值，而是访问变量中的内存地址，只要指针数据变化对应的变量值会跟随变化
	C := &A
	fmt.Printf("%T %d %p %v\n", C, *C,C,C)   //%p打印指针值（内存地址），%v适合不清楚类型时打印值使用
	*C = 3
	fmt.Printf("%d\n", A)

}
