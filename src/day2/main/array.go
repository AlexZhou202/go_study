package main

import "fmt"

func main() {
	/*
		数组是具有相同数据类型的数据项组成的一组长度固定的序列，数据项叫做数组的元素，数组的长度必须是非负整数的常量，长度也是类型的一部分
		注意：
	*/
	var nums [10] int
	var t2 [5] bool
	var t3 [3] string
	fmt.Printf("%T %d\n", nums, nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)

	//字面量
	nums = [10] int{10, 20, 30}
	fmt.Printf("%d\n", nums)
	nums = [10] int{0: 10, len(nums) - 1: 20}
	fmt.Printf("%d %d\n", nums, len(nums))
	//简短声明，注意：如果定义了数组长度，数组元素长度必须与定义的数组长度相同
	numb := [...]int{1, 2, 3}
	fmt.Println(numb)
	numc := [10]int{1, 2, 3}
	fmt.Println(numc)
	numd := [10]int{2, 4, 6}

	//操作，数组长度不相同不能进行操作
	fmt.Println(numc == numd)

	//获取数组长度
	fmt.Println(len(numd))

	//索引
	fmt.Println(numd[0], numd[1])
	numd[0], numd[2] = 1000, 20
	numd[4] = 200
	fmt.Println(numd)

	for i := 0; i < len(numd); i++ {
		fmt.Println(i, ":", numd[i])
	}

	for _, v := range numd {
		fmt.Println(v)
	}

	//切片
	fmt.Printf("%T %d\n", numd[0:4], numd[0:4])
	fmt.Printf("%T %d\n", numd[0:4:6], numd[0:4:6])

}
