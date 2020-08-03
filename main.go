package main

import (
	"fmt"

	"github.com/gritt/hello"
	hello2 "github.com/gritt/hello/v2"
)

func main() {
	fmt.Println(hello.HelloWorld())

	msg, l := hello2.HelloWorld()
	fmt.Printf("%s  has %d chars", msg, l)
}
