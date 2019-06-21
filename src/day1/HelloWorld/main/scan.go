package main

import "fmt"

func main() {
	var v_input string
	fmt.Println("请输入名字")
	fmt.Scan(&v_input)
	fmt.Printf("你输入的名字是\t%s\n",v_input)

	var v_age int
	fmt.Println("请输入年龄")
	fmt.Scan(&v_age)
	fmt.Printf("你输入的年龄是\t%d\n",v_age)


	var v_height float64
	fmt.Println("请输入身高")
	fmt.Scan(&v_height)
	fmt.Printf("你的身高是\t%.2f\n",v_height)


}
