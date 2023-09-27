package main

import "fmt"

func exibirSinal(n int) {
	if n < 0 {
		fmt.Println("N")
	} else {
		fmt.Println("P")
	}
}

func main() {
	numeroneg := -5
	numeropos := 5
	numeroull := 0
	exibirSinal(numeroneg)
	exibirSinal(numeropos)
	exibirSinal(numeroull)
}
