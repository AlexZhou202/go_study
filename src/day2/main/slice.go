package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%T %q\n", nums, nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) //cap(nums) 获取切片容量
	fmt.Printf("%v\n", nums == nil)                       //切片声明需要指定组成元素的类型(实际是指针指向为nil)，但不需要指定存储元素的数量（长度）。在切片声明后，会被初始化为 nil，表示暂不存在的切片

	//字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)

	//数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4, 5, 6}
	nums = arrays[1:10]
	fmt.Printf("%#v\n", nums)

	//make函数
	nums = make([]int, 3) //初始化切片，默认值为0
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = make([]int, 3, 5) //初始化切片，长度3，容量5
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	//元素操作(增，删，改，查)
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	nums[2], nums[0] = 100, 200
	fmt.Println(nums)
	nums = append(nums, 1) //如果长度设置为3，容量设置为5，不能直接访问长度为4,5的元素，需要使用append插入元素
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 22) //如果长度设置为3，容量设置为5，不能直接访问长度为4,5的元素，需要使用append插入元素
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 33) //当容量不足时，使用append插入元素后容量会自动扩展(翻倍扩展(1-1.5倍之间))
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	//遍历切片
	for _, i := range nums {
		fmt.Println(i)
	}
	//切片操作
	fmt.Printf("%#v\n", nums[1:5])
	//容量:n_cap - start(10-1)
	nums = make([]int, 3, 10)
	n := nums[1:3:10]
	fmt.Printf("%#v %d %d\n", n, len(n), cap(n))
	//容量:scr_cap - start(10-2)
	n = nums[2:3]
	fmt.Printf("%#v %d %d\n", n, len(n), cap(n))

	//直接赋值出现的问题，赋值是引用原切片的引用，所以当被赋值的切片更改元素时原切片的值也会跟着改变，对于该问题使用复制操作可以解决
	n[0] = 2
	fmt.Println(nums, n) //[0 0 2] [2]
	//因n引用的是nums的引用，所以只要当原切片新增元素与被赋值的切片元素相同时，被赋值的切片元素值就会被覆盖，该问题会持续到超出容量后才会进行重新映射解决
	n = append(n, 5)
	fmt.Println(nums, n, len(nums), len(n)) //[0 0 2] [2 5] 3 2
	nums = append(nums, 10)
	fmt.Println(nums, n, len(nums), len(n)) //[0 0 2 10] [2 10] 4 2

	//复制操作，注意：当两个切片长度不一样的时候按照被赋值的切片元素最大长度的元素进行复制
	numa := []string{"aaa", "bbb"}
	numb := []string{"ddd", "ccc"}
	copy(numa, numb) //将numb的元素复制到numa中
	fmt.Printf("%#v %s %s\n", numa, numa, numb)
	numa[0] = "eee"
	fmt.Printf("%#v %s %s\n", numa, numa, numb)

	//移除操作
	elements := []int{1, 2, 3, 4, 5}
	//copy(elements[3:],elements[4:]) //将最后一个元素赋值给上一个元素
	copy(elements[len(elements)-2:], elements[len(elements)-1:])                               //将最后一个元素赋值给上一个元素
	fmt.Println(elements[:len(elements)-1], elements, len(elements), elements[:len(elements)]) //使用切片将最后一个元素切除(切片包头不包尾，所以[:4]实际只会输出4个元素)

	//队列操作，每次都获取头一个元素，先进先出
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 3)
	queue = append(queue, 2)
	fmt.Println(queue[0], queue) //输出第一个元素,切片值:[1 3 2]
	queue = queue[1:]            //切片获取第二个元素,切片值:[3 2]
	fmt.Println(queue[0], queue) //输出第二个元素
	queue = queue[1:]            //切片获取第三个元素,切片值:[2]
	fmt.Println(queue[0], queue) //输出第三个元素

	//堆栈操作，每次都获取最后一个元素，先进后出
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	fmt.Println(stack[len(stack)-1], stack) //输出最后一个元素,切片值[1 3 2]
	stack = stack[:len(stack)-1]            //获取倒数第二个元素,切片值[1 3]
	fmt.Println(stack[len(stack)-1], stack) //输出倒数第二个元素
	stack = stack[:len(stack)-1]            //获取第一个元素,切片值[1]
	fmt.Println(stack[len(stack)-1], stack) //输出第一个元素

	//数组是值类型，当一个数组对另外一个数组赋值时，是会使用另一个内存空间去存储的，所以不会出现引用的问题
	arrays01 := [3]int{1, 3, 2}
	arrays02 := arrays01
	arrays02[0] = 20
	fmt.Println(arrays01, arrays02) //[1 3 2] [20 3 2]

}
