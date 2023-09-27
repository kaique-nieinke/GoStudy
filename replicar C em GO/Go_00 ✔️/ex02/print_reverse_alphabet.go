package main

import "fmt"

func alphavetReverse() {
	alfabeto := ""
	for i := 'z'; i >= 'a'; i-- {
		alfabeto += string(i)
	}
	fmt.Println(alfabeto)
}

func main() {
	alphavetReverse()
}
