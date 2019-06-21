package main

import "fmt"

func main() {
	//"" =>可解释的字符
	// `` => 原生字符串
	// 特殊字符 \r \n \t \f \b \v
	var name string = "k\tk"
	var name2 string = `k\tk`
	var name3 string = "k\\tk"
	fmt.Println(name,name2,name3)

	//操作
	//算术运算符: + (连接)
	fmt.Println("我叫"+"kk")
	//关系运算符(== != > >= < <=)
	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" < "bb")
	fmt.Println("ab" <= "bbb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" >= "bb")
	fmt.Println("bb" > "b")
	//赋值
	s := "我叫"
	s += "zjj"
	fmt.Println(s)

	//字符串定义内容必须只能为ascii（不为中文unicode）
	//索引 0 n-1 (n字符长度)
	desc := "abcdef"
	fmt.Printf("%T %c\n",desc,desc[0])
	//切片[start:end] start end -1
	fmt.Printf("%T %s\n",desc[0:2],desc[0:2])
	//获取字符长度
	fmt.Println(len(desc))
}
