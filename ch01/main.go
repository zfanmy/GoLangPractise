package main

import (
	"fmt"

	"zfanmy.com/gopl/ch01/dup"
	"zfanmy.com/gopl/ch01/helloworld"
)

func main() {
	helloworld.HelloWorld()
	//echo.PrintEcho()
	//dup.Dup()
	dup.TestStableUseSlice()
	fmt.Println("//////////////////////////////////////////////////////////")
	dup.TestStableUsePair()
	fmt.Println("//////////////////////////////////////////////////////////")
	dup.TestStandardStable()
}
