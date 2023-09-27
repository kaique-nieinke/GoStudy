package main

import "fmt"

func main() {
	numeros := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}

	soma1 := 0
	soma2 := 0

	for i := 0; i < len(numeros); i++ {
		if i < 5 {
			soma1 += numeros[i]
		} else {
			soma2 += numeros[i]
		}
	}

	fmt.Println("Soma dos números de 1 a 5:", soma1)
	fmt.Println("Soma dos números de 6 a 10:", soma2)
}
