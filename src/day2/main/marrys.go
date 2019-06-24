package main

import "fmt"

func main() {
	var marrys [3][2]int
	//长度为2的int类型数组 => [2]int
	fmt.Println(marrys)
	fmt.Println(marrys[0])
	fmt.Println(marrys[0][0])

	marrys[0] = [2]int{1, 3} //将元素0的数组元素修改成1,3
	fmt.Println(marrys)

	marrys[1][1] = 1000 //将数组1里的数组1的数组元素修改成1000
	fmt.Println(marrys)

	marrys = [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(marrys)

	//二维数组遍历
	for i, _ := range marrys { //先获取索引值
		for _, j := range marrys[i] { //利用索引值获取子数组的值
			fmt.Println(j)
		}
	}

	for _, v := range marrys { //先获取索引值
		for _, j := range v { //利用索引值获取子数组的值
			fmt.Println(j)
		}
	}
}
