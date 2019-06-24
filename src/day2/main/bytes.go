package main

import (
	"bytes"
	"fmt"
)

func main() {
	bytesa:=[]byte{'1','2','3'}
	fmt.Printf("%#v\n",bytesa)
	bytesa_s:=string(bytesa) //将字节切片转换成字符串
	fmt.Printf("%#v\n",bytesa_s)
	bytesb:=[]byte(bytesa_s) //将字符串转换成字节切片
	fmt.Printf("%#v\n",bytesb)

	//字节切片比较，与strings包相同
	fmt.Println(bytes.Compare([]byte("abcd"),[]byte("bcde")))
	fmt.Println(bytes.Index([]byte("eddd"),[]byte("e")))
	fmt.Println(bytes.Contains([]byte("eddd"),[]byte("e")))
}
