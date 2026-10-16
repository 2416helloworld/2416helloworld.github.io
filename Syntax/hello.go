package main

/*
	导入的包未使用则会报错
*/
import (
	"fmt"
	"time"
)

func main() { //函数的{一定要和函数名在同一行
	fmt.Println("Hello World") //一条语句可以不加;

	time.Sleep(1 * time.Second)
}
