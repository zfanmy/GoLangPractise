package echo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PrintEcho 输出命令行参数
func PrintEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// 使用循环在字符串slice中添加间隔符
func echo() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 练习1.1
func echo1() {
	fmt.Println(os.Args[0])
}

// 练习1.2
func echo2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(index+1) + " " + arg)
	}
}
