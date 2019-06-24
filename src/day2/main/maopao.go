package main

import "fmt"

func main() {
	//冒泡排序
	heght := []int{10, 6, 7, 9, 5}
	//遍历
	for n := 0; n < len(heght)-1; n++ { //重复内部元素比较次数
		for i := 0; i < len(heght)-1; i++ { //内部元素比较，逐个元素进行比较，将最大的数移动到最后
			if heght[i] > heght[i+1] {
				heght[i], heght[i+1] = heght[i+1], heght[i]
			}
		}
	}
	fmt.Println(heght)
}
