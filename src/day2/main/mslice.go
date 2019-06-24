package main

import "fmt"

func main() {
	//声明及初始化
	nums := [][]int{{1, 2}, {1, 3}, {1, 4}}
	fmt.Printf("%#v %v %v", nums, nums[0], nums[0][0])

	//修改操作
	nums[0] = []int{2, 100}
	nums[1][1] = 200
	fmt.Println(nums[0], nums[1][1], nums)

	//切片
	fmt.Println(nums[0:2])
	fmt.Println(nums[0 : len(nums)-1])

	//遍历
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Printf("%v\n", nums[i][j])
		}
	}

	for _, v := range nums {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//append
	nums = append(nums, []int{2, 100, 1}) //直接添加一个子切片
	fmt.Println(nums)
	nums[0] = append(nums[0], 30) //直接在原本的切片中添加元素
	fmt.Println(nums)

	//copy，注意拷贝对象子切片
	nums2 := [][]int{{}}
	copy(nums2, nums)
	fmt.Println(nums2) //[[2 100 30]]
	nums2 = [][]int{{}, {}, {}}
	copy(nums2, nums)
	fmt.Println(nums2) //[[2 100 30] [1 200] [1 4]]
}
