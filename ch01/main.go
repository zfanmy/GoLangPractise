package main

import (
	"fmt"

	"zfanmy.com/gopl/ch01/dup"
	"zfanmy.com/gopl/ch01/echo"
	"zfanmy.com/gopl/ch01/fetch"
	"zfanmy.com/gopl/ch01/helloworld"
	"zfanmy.com/gopl/ch01/lissajous"
)

func main() {
	// 1.1 Hello World!
	helloworld.HelloWorld()
	fmt.Println("//////////////////////////////////////////////////////////")
	// 1.2 输出命令行参数
	echo.PrintEcho()
	fmt.Println("//////////////////////////////////////////////////////////")
	// 1.3 找出重复行
	dup.Dup()
	dup.TestStableUseSlice() // 测试1.18 sort.SliceStable() 获取稳定的map排序
	fmt.Println("//////////////////////////////////////////////////////////")
	dup.TestStableUseSlice() // 重复执行，验证排序稳定性
	fmt.Println("//////////////////////////////////////////////////////////")
	// 1.4 GIF 动画
	lissajous.Paint()
	fmt.Println("//////////////////////////////////////////////////////////")
	// 1.5 获取URL内容
	fetch.GetURL()

}
