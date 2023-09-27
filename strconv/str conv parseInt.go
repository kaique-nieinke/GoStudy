package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Printf("%T \n", b)
	fmt.Println(b)
}
