package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{}
	num = append(num, 3)
	num = append(num, 5)
	num = append(num, 1)
	num = append(num, 2)
	fmt.Println(num)
	sort.Ints(num) //排序
	fmt.Println(num)

	name2 := []string{"bb", "cc", "aa"}
	sort.Strings(name2)
	fmt.Println(name2)

	name3 := []float64{3.1, 1.3, 1.2, 3.2}
	sort.Float64s(name3)
	fmt.Println(name3)

}
