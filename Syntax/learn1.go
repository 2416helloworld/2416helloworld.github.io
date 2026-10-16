package main

/*
	声明变量的方式
	方式1，2，3可以用来声明全局变量
	而方式4只能用于函数体声明
*/
import (
	"fmt"
)

func main() {
	//方式1：声明一个变量，默认值为0
	var a int
	fmt.Println(a)
	//方式2：声明一个变量，初始化为一个值
	var b int = 10
	fmt.Println(b)
	//方式3：在初始化时，可以省去数据类型，通过自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println(c)
	//方式4：可以省去var，直接自动匹配
	d := 10

	fmt.Println(d)
}
