package main

import "fmt"

func main() {
	users :=[]string{"zjj","jack","jack","zjj","cls"}
	//定义映射
	user_map:=map[string]int{}
	for _,v :=range users{
		//fmt.Println(user_map)
		user_map[v]++ //根据int返回零值的特点，如获取到key，则将value值+1，等价于user_map[v] = user_map[v] + 1
	}
	fmt.Println(user_map)
}
