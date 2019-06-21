package main

import "fmt"

func main() {
	//goto,定义lable进行跳转，一般用于for循环很深后配搭break结束整个循环时使用
	//使用goto写1..100的和
	v_sum := 0
	i := 0
START:
	if i > 100 {
		goto FOREND
	}
	v_sum += i
	i++
	goto START

FOREND:
	fmt.Println(v_sum)

	END:
for i:=1;i<=5;i++{
	for j:=1;j<=5;j++{
		if i*j ==9{
			break END
		}
		fmt.Println(i,j,i*j)
	}
}


}
