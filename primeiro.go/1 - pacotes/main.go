package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("kaiquenieinke@gmail.com")
	fmt.Println(erro)
}
