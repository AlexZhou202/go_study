package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abcd","bcd")) //比较两个字符是否相等
	fmt.Println(strings.Contains("abcdefg","a")) //判断后面的字符是否在前面的字符串中包含
	fmt.Println(strings.ContainsAny("abcde","abc")) //判断后面的任意字符是否在前面的字符串中包含
	fmt.Println(strings.Count("abcdefgabd","a")) //判断后面的字符在前面的字符串中出现多少次
	fmt.Printf("%q\n",strings.Fields("abcdefgabd aaa\nddd\twwww\vddda")) //按照空格符分隔，包括(\n,\t,\v,\r,\f)
	fmt.Println(strings.HasPrefix("abcdefgabd","a")) //判断后面的字符是否在前面的字符串中作为前缀
	fmt.Println(strings.HasSuffix("abcdefgabd","a")) //判断后面的字符是否在前面的字符串中作为后缀
	fmt.Println(strings.Index("abcdefgabd","b")) //判断后面的字符在前面的字符串中查找的首次索引位置
	fmt.Println(strings.LastIndex("abcdefgabd","d")) //判断后面的字符在前面的字符串中查找的最后索引位置
	fmt.Printf("%q\n",strings.Split("abcdefgabd;aaa;ddd;wwww;ddda",";")) //以特定字符进行分隔
	fmt.Printf("%q\n",strings.Join([]string{"abcd","bdc","ddd"},":")) //使用特定字符对切片进行分隔
	fmt.Println(strings.Repeat("abc",3)) //重复输出N次字符
	fmt.Println(strings.Replace("abcdbadada","abc","xxx",1)) //将最前的字符串中的字符替换成新的字符，n表示替换多少次
	fmt.Println(strings.Replace("abcdbadada","abc","xxx",-1)) //等价于replaceall，全部替换的意思
	fmt.Println(strings.ReplaceAll("abcdbadada","abc","xxx")) //将最前的字符串中的字符替换成新的字符，全部替换
	fmt.Println(strings.ToLower("ABCDSDS")) //将字符串替换成小写
	fmt.Println(strings.ToUpper("addasdsa")) //将字符串替换成大写
	fmt.Println(strings.ToTitle("aDSDSa")) //首字母大写
	fmt.Println(strings.Trim("ABCDSDS","AB")) //将指定的字符从字符串中删除掉
	fmt.Println(strings.TrimSpace("   aaa dddbbcccc\n \t")) //去掉字符串首尾空白字符(\n,\t,\v,\r,\f)

}
