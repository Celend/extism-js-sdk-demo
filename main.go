package main

import (
	"fmt"
	"github.com/extism/go-pdk"
)

func main() {}

//export hello
func hello() int32 {
	fmt.Println("before:", pdk.InputString())
	pdk.OutputString("Hello World")
	fmt.Println("after")
	return 0
}
