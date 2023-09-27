package main

import (
	"fmt"
	"golangestudo/model"
)

func main() {
	fmt.Println("iniciando..")

	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "XPTO",
		Modelo: "CG",
	}
	Moto := modelMoto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}
	fmt.Println(Moto)
}
