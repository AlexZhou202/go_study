package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		yes string
	)
	fmt.Println("有卖西瓜的吗?")
	fmt.Scan(&yes)
	println(strings.ToUpper(yes)) //将变量值转换为大写
	fmt.Println("老婆的想法!")
	if yes == "y" || yes == "Y" {
		fmt.Println("买一个西瓜")
	} else {
		fmt.Println("买十个包子")
	}
	fmt.Println("有卖西瓜的吗?")
	fmt.Scan(&yes)
	fmt.Println("老公的想法!")
	if yes == "y" || yes == "Y" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买十个包子")
	}

	var (
		score int
	)
	fmt.Println("请输入成绩")
	fmt.Scan(&score)
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

}
