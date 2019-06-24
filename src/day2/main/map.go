package main

import "fmt"

func main() {
	//映射是存储一系列无序的key/value对，通过key来对value进行操作（增、删、改、查）。
	//映射的key只能为可使用==运算符的值类型（字符串、数字、布尔、数组），value可以为任意类型s
	var nums map[string]int //定义一个key为字符value为整数的空映射
	fmt.Printf("%#v\n", nums)
	user_map:=map[string]int{}
	fmt.Println(user_map)
	nums = make(map[string]int) //使用make函数初始化
	fmt.Println(nums)
	fmt.Println(nums == nil)
	//字面量
	nums = map[string]int{"zjj": 8, "jk": 9, "jack": 10}
	fmt.Println(nums)

	//增、删、改、查
	fmt.Println(nums["zjj"]) //查询key为zjj的value值
	if v, ok := nums["jack"]; ok == true { //对于没有查询的值，会返回value类型的零值,所以可以通过if进行判断映射的key是否存在，如存在则输出value值v
		fmt.Println(v)
	}

	//增加及修改
	nums["jk"] = 20 //修改映射的值，查询原来的key对value进行修改
	fmt.Println(nums)
	nums["ww"] = 30 //增加映射的key及value值，key不存在时则为新增
	fmt.Println(nums)

	//删除
	delete(nums, "ww") //删除映射对应的key值
	fmt.Println(nums)

	//获取元素的长度
	fmt.Println(len(nums))

	//遍历映射
	for k, v := range nums {
		fmt.Println(k, v)
	}

	//key至少可以有== != 运算，key值：bool，整数，字符串，数组
	//value => 任意类型
	var name = map[string][]string{"zjj": {"gz", "15018773825", "60"}, "boy": {"dd", "13902221321", "100"}} //定义一个key为姓名，value为切片的映射
	fmt.Println(name)
	var name2 = map[string]map[string]string{"zjj": {"地方": "gz", "联系方式": "150187722xxx", "成绩": "100"}} //定义一个key为姓名,value为映射的映射
	fmt.Println(name2)

	//判断key是否存在
	if v, ok := name2["zjj"]; ok == true {
		fmt.Println(v, ok)
	}
	name2["zjj"] = map[string]string{"地方": "gd", "联系方式": "222", "成绩": "200"} //修改映射key的value值
	fmt.Println(name2)
	name2["zjj"]["地方"] = "haizhu" //修改value里对应映射的值
	fmt.Println(name2)
	name2["ww"] = map[string]string{"地方": "gd", "联系方式": "222", "成绩": "200"} //添加映射key的value值
	fmt.Println(name2)
	delete(name2["ww"], "成绩") //删除映射key对应的value映射key
	fmt.Println(name2)
	delete(name2, "ww") //删除映射key ww
	fmt.Println(name2)

}
