//定义main包
package main
//导入fmt,time包
import (
	"fmt"
	"time"
)

/*
main函数
程序启动的入口
 */

func main() {
	fmt.Println("HelloWorld")  //输出helloworld
	time.Sleep(5 * time.Second) //加了等待5秒防止windows编译成exe双击一闪而过的问题
}
