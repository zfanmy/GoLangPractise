package echo

import (
	"fmt"
	"strings"
	"testing"
)

//循环与使用strings.Join基准测试对比

var testData = []string{"a", "b", "c", "d", "e", "f"}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, data := range testData {
			s += sep + " " + data
			sep = ""
		}
		//fmt.Println(s)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintln(strings.Join(testData, " "))
	}
}
