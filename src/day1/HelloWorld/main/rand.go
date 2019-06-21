package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())  //如果不添加rand则不会随机生成数据
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Int() % 100)


}
