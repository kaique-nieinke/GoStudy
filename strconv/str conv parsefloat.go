package main

import (
	"fmt"
	"strconv"
)

func main() {
	texto := "42.55"
	b, _ := strconv.ParseFloat(texto, 64)
	fmt.Printf("%T \n", b)
	fmt.Println(b)
}
