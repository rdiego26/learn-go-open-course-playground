package main

import (
	"fmt"
)

type data int

var x data
var y int

func main() {
	fmt.Printf("x=%v\n", x)
	fmt.Printf("x type is %T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(x)
	fmt.Println(y)
}
