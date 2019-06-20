package main

import "fmt"

func main() {
	//作用域：定义标识符可以使用的范围
	//在GO中使用{}来定义作用域的范围
	//使用原则：子语句块可以使用父语句块的标识符，而父不能使用子语句块的标识符
	outer := 1   //子作用域可以对父作用域重新进行定义
	{
		inner:=2
		fmt.Println(outer)
		fmt.Println(inner)
		outer:=3
		{
			inner2:=3
			inner:=10
			fmt.Println(inner)
			fmt.Println(inner2)
			fmt.Println(outer)
		}
	}
}

