package main

import "fmt"

func main() {
	const name string = "zjj" //常量赋值
	fmt.Println(name)

	const n ,m string = "aa","wwaa" //常量相同类型赋值

	fmt.Println(n,m)

	const (	 //常量多个类型赋值
		abc int = 10
		dba string = "dba"
	)
	fmt.Println(abc,dba)

	const dd = "zjj"	//常量自动推导赋值
	fmt.Println(dd)

	const (		//常量作用域赋值
		ww = 10
		ee = 1.37
	)
	fmt.Println(ww,ee)

	const qq,gg,ff = "ww","dd",1.38	//多个常量自动推导
	fmt.Println(qq,gg,ff)

	const (			//对于赋值，如果后边的变量什么都不写，则会赋予第一个含有赋予的值
		C7 int = 1
		C8
		C9
	)

	fmt.Println(C7,C8,C9)

	//枚举，const+iota (每个括号都会重新计数)
	const(
		E1 int = iota
		E2
		E3
	)

	fmt.Println(E1,E2,E3)

}