package main

import (
	"fmt"
)

func main() {
	var listatoda = []int{2, 10, 9, 4, 5, 8, 1, 3, 2, 11}
	segundaLista := listatoda[:3]
	terceiraLista := listatoda[4:]
	ultimoItem := listatoda[len(listatoda)-1:]

	fmt.Println(segundaLista)
	fmt.Println(terceiraLista)
	fmt.Println(ultimoItem)

}
