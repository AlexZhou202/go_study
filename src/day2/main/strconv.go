package main

import (
	"fmt"
	"strconv"
)

func main() {

	//将字符串转换成bool
	if v,err:=strconv.ParseBool("true");err==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}


	//将字符串转换成int
	if v,err:=strconv.Atoi("222231");err==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}

	//将字符串转换成int，s:字符串 base:进制(10进制，16进制，8进制)，bitsize：长度/范围
	if v,err:=strconv.ParseInt("222231",8,64);err==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}


	//将字符串转换成float，s:字符串 bitsize：长度/范围
	if v,err:=strconv.ParseFloat("1.234",64);err==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}

	//将其他类型转换成字符串
	sd:=fmt.Sprintf("%d",123)
	fmt.Println(sd)

	sf:=fmt.Sprintf("%.3f",1.321)
	fmt.Println(sf)

	fmt.Printf("%q\n",strconv.Itoa(12311))
	fmt.Printf("%q\n",strconv.FormatBool(false))
	fmt.Printf("%q\n",strconv.FormatInt(111,16)) //将int转换成字符串，i:int base:进制(10进制，16进制，8进制)
	/*
		/ FormatFloat 将浮点数 f 转换为字符串值
		// f：要转换的浮点数
		// fmt：格式标记（b、e、E、f、g、G）
		// prec：精度（数字部分的长度，不包括指数部分）
		// bitSize：指定浮点类型（32:float32、64:float64）
		//
		// 格式标记：
		// 'b' (-ddddp±ddd，二进制指数)
		// 'e' (-d.dddde±dd，十进制指数)
		// 'E' (-d.ddddE±dd，十进制指数)
		// 'f' (-ddd.dddd，没有指数)
		// 'g' ('e':大指数，'f':其它情况)
		// 'G' ('E':大指数，'f':其它情况)
		//
		// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
		// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	*/
	fmt.Printf("%q\n",strconv.FormatFloat(1.23,'f',3,64))




}
