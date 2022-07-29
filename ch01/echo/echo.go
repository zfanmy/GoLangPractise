package echo

import (
	"fmt"
	"os"
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
