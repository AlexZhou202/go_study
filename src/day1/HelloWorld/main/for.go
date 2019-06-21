package main

import "fmt"

func main() {
	//1+100之和
	var (
		v_sum int
	)
	for i := 1; i <= 100; i++ {
		v_sum += i
	}
	fmt.Printf("1..100的结果为:\t%d\n", v_sum)

	v_sum = 0
	i := 1
	for i <= 100 {
		v_sum += i
		i++
	} //while写法
	fmt.Println(v_sum)

	i = 0
	for { //死循环写法，等价于while true
		fmt.Println(i)
		i++
	}

	//字符串，数组，切片，映射，管道
	desc := "我爱中国"
	for i, y := range desc {
		//fmt.Println(i, y)
		fmt.Printf("%d %T %q\n",i,y,y)
	}

}
