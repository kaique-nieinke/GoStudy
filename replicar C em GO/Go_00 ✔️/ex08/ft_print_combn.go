package main

import "fmt"

func exibirCombinacoes(n int) {
	for i := 0; i < 10-n; i++ {
		for j := i + 1; j < 10; j++ {
			fmt.Printf("%d%d", i, j)
			if i != 8 || j != 9 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}

func main() {
	n := 2
	fmt.Printf("Combinações de %d números em ordem crescente:\n", n)
	exibirCombinacoes(n)
}
