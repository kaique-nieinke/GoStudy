package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, _ := strconv.ParseBool("true")
	fmt.Printf("%T \n", b)
	fmt.Println(b)
}
