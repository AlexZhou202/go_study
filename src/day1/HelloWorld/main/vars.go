package main

import "fmt"

var aa string = "test"

func main() {
	/*
		定义一个string类型的变量me
		变量名的命名规则命名规则：
		a. 只能由非空字母(Unicode)、数字、下划线(_)组成
		b. 只能以字母或下划线开头
		c. 不能Go语言关键字(25个)
		d. 避免使用Go语言预定义标识符,true/false/nil/bool/iota
		e. 建议使用驼峰式
		f. 标识符区分大小写
		注意：函数外定义变量不受GO引用变量的检查
	*/

	var me string = "zjj"
	fmt.Println(me)
	fmt.Println(aa)

	var name, user string = "aa", "bb" //多个变量同类型定义
	fmt.Println(name, user)

	var ( //括号赋值变量定义  常用
		age    int     = 10
		height float64 = 1.83
	)

	fmt.Println(age, height)

	var ( //括号赋值自动推导    常用
		dd = 10
		ee = "iou"
	)
	fmt.Println(dd, ee)

	var ss, aa = "kk", 1.38 //自动推导
	fmt.Println(ss, aa)

	isBoy := true //简短声明(可以不使用var进行定义，由原来的=改变为:=)，注意只能在函数内部使用   常用
	fmt.Println(isBoy)

	ss, aa, isBoy = "ddd", 5.3, false //多个变量自动推导
	fmt.Println(ss, aa, isBoy)

	ff, gg := aa, ss //变量值交换
	fmt.Println(ff, gg)

}
