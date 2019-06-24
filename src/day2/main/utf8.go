package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ch:="我爱中国人民共和国"
	fmt.Println(len(ch)) //输出原始字节长度
	fmt.Println(utf8.RuneCountInString(ch)) //输出unicode码的长度
}
