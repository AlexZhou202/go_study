package main

import "fmt"

func main() {
	//整数类型
	//标识符:int/int*/uint/uint*/uintptr/byte/rune
	//字面量:十进制，八进制0777，十六进制0X10
	var age int = 31
	fmt.Printf("%T %d\n", age, age)
	fmt.Println(0777, 0x10)

	//操作
	//算术运算（+,-,*,/,%）
	fmt.Println(1 + 2)
	fmt.Println(3 - 10)
	fmt.Println(3 * 9)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)

	//自增，自减
	age++ //自增
	fmt.Println(age)
	age-- //自减
	fmt.Println(age)

	//关系运算(== != > >= < <=)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 > 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 < 3)
	fmt.Println(2 <= 3)

	//位运算 二进制运算 10 =>2
	// & | ^ << >> &^    &与运算，或运算，异或运算，左移，右移，双目运算
	// 2 = 0010 7=0111
	//7 & 2 => 0111 & 0010 = 0010 => 2   有0则0
	//7 | 2 => 0111 | 0010 = 0111 => 7   有1则1
	//7 ^ 2 => 0111 ^ 0010 = 0101 => 5   相同为0，相异为1
	//2 << 1 => 0010 << 1 => 0100 => 4   左移N位
	//2 >> 1 => 0010 >> 1 => 0001 => 1   右移N位
	//7 &^ 2 => 0111 &^ 0010 = 0101 => 5  相同为0，不同保留
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(7 ^ 2)
	fmt.Println(2 << 1)
	fmt.Println(2 >> 1)
	fmt.Println(7 &^ 2)

	//赋值（=,+=,-=,*=,/=,%=,&=,|=,^=,<<=,>>=,&^=）
	age = 1
	age += 3
	fmt.Println(age)
	age -= 2
	fmt.Println(age)
	age *= 2
	fmt.Println(age)
	age /= 4
	fmt.Println(age)

	//类型转换，不同的数据类型是不能进行运算，所以需要进行类型转换
	var intA int = 10
	var intB uint = 20
	fmt.Println(intA + int(intB))
	fmt.Println(uint(intA) + intB)
	//大 -> 小 转换会出现溢出问题，一般都是从小往大转换
	var intC int = 0xFFFF
	fmt.Println(intC, uint8(intC), int8(intC)) //uint8溢出输出255，int8溢出输出-1

	//byte,rune
	var aa byte = 'A' //输出整数结果(ascii码)65
	var dd rune = '我' //输出整数结果25105
	fmt.Println(aa, dd)
	fmt.Printf("%T %d %b %o %x \n", age, age, age, age, age) //%d 整数 %b 二进制 %o 八进制 %x 十六进制
	fmt.Printf("%T %c \n", aa, aa)                           //%c byte字符
	fmt.Printf("%T %q %U \n", dd, dd, dd)                    //%q unicode字符 %U unicode码
	fmt.Printf("%d %d \n", aa, dd)
	//格式化转换
	fmt.Printf("%5d %d \n", aa, dd)  //%5d长度占5位，从头占位
	fmt.Printf("%05d %d \n", aa, dd) //%5d长度占5位，不足往左补0
	fmt.Printf("%-5d %d \n", aa, dd) //%5d长度占5位，从尾占位

	//float32,float64
	//字面量:十进制表示法 科学技术表示法
	//M E N => M * 10 ^ N
	//1.9E-1 => 1.9 * 10^-1 = > 0.19
	var heghit float64 = 1.68
	fmt.Printf("%T %f\n", heghit, heghit) //%f 小数
	var weight float64 = 13.05E1
	fmt.Printf("%T %f\n", weight, weight)
	fmt.Println(weight)

	//浮点型算术运算符(+,-,*,/,++.--)，小数进行算术运算时有可能会出现精度损耗（自增自减肯定会出现）
	fmt.Println(1.1 + 1.2)
	fmt.Println(1.1 - 1.2)
	fmt.Println(1.1 * 1.2)
	fmt.Println(1.1 / 1.2)
	heghit++ //
	fmt.Println(heghit)
	heghit-- //小数进行自增自减时会出现精度损耗
	fmt.Println(heghit)
}
