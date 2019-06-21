package main

import "fmt"

func main() {
	var (
		yes string
		score int
	)
	fmt.Println("有卖西瓜的吗")
	fmt.Scan(&yes)
	switch yes{
	case "y","Y":fmt.Println("买一个西瓜")
	default:
		fmt.Println("买十个包子")
		
	}
	fmt.Println("请输入成绩:")
	fmt.Scan(&score)
	switch {
	case score>=90 && score<=100:
		fmt.Println("A")
	case score>=80 && score<90:
		fmt.Println("B")
	case score>=70 && score<80:
		fmt.Println("C")
	case score>=60 && score<70:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
