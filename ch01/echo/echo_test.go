package echo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

//循环与使用strings.Join基准测试对比

var testData = []string{"a", "b", "c", "d", "e", "f"}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ":" + testData[0]
		s += ":" + testData[1]
		s += ":" + testData[2]
		s += ":" + testData[3]
		s += ":" + testData[4]
		s += ":" + testData[5]
		_ = s
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3],
			testData[4], testData[5])
		_ = s
	}
}

func BenchmarkBufferString(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(testData[0])
		buf.WriteString(":")
		buf.WriteString(testData[1])
		buf.WriteString(":")
		buf.WriteString(testData[2])
		buf.WriteString(":")
		buf.WriteString(testData[3])
		buf.WriteString(":")
		buf.WriteString(testData[4])
		buf.WriteString(":")
		buf.WriteString(testData[5])
		s := buf.String()
		_ = s
	}
}

func BenchmarkBufferByte(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.WriteString(testData[0])
		buf.WriteByte(':')
		buf.WriteString(testData[1])
		buf.WriteByte(':')
		buf.WriteString(testData[2])
		buf.WriteByte(':')
		buf.WriteString(testData[3])
		buf.WriteByte(':')
		buf.WriteString(testData[4])
		buf.WriteByte(':')
		buf.WriteString(testData[5])
		s := buf.String()
		_ = s
	}
}

/*
测试结果
strings.Join
24049521
48.14 ns/op
16 B/op
1 allocs/op
VirtualApple @ 2.50GHz
concat
6540340
185.4 ns/op
48 B/op
5 allocs/op
VirtualApple @ 2.50GHz
Sprintf
4010121
298.8 ns/op
112 B/op
7 allocs/op
VirtualApple @ 2.50GHz
BufferString
13136036
108.0 ns/op
0 B/op
0 allocs/op
VirtualApple @ 2.50GHz
BufferByte
19703556
98.07 ns/op
0 B/op
0 allocs/op
VirtualApple @ 2.50GHz
*/
