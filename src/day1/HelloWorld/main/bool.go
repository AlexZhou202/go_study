package main

import "fmt"

func main() {
	/*
		布尔类型 表示真假
		标识符bool
		字面量(可选值)true/false
		零值(默认值) false
	*/
	var zero bool //零值为false
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
	//操作
	//逻辑运算(与&&,或||,非!)
	//aBool && bBool 当两个为true的时候才是true
	//aBool || bBool 当其中一个为true时就是true
	//!(aBool &&或|| bBool) 与跟或的结果取反
	//关系运算(==,!=)
	fmt.Println("&&运算操作")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false && true)

	fmt.Println("||运算操作")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(false || true)

	fmt.Println("!运算操作")
	fmt.Println(!(true && true))
	fmt.Println(!(true && false))
	fmt.Println(!(false && false))
	fmt.Println(!(false && true))

	//关系运算符(==,!=)
	fmt.Println("关系运算操作")
	fmt.Println(isGirl == isBoy)
	fmt.Println(isBoy != zero)
	fmt.Println("bool类型打印")
	fmt.Printf("%T %t %t\n", isBoy, isBoy, isGirl) //占位符%t代表bool输出
}
