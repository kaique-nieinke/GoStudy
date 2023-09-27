package main

import "fmt"

func todosOsDigitos() {
	for i := 1; i <= 9; i++ {
		fmt.Print(i, "")
	}

	fmt.Println()
}

func main() {
	todosOsDigitos()
}
