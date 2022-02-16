package main

import (
	"fmt"
)

var x, y, z = 42, "James Bond", true

func main() {
	s := fmt.Sprintf("%v - %v - %v", x, y, z)
	fmt.Println(s)
}
