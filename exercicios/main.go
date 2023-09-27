package main

import (
	"ex01/model"
	"fmt"
	"time"
)

func main() {

	var nomeDoItens []string
	nomeDoItens = append(nomeDoItens, "arroz")
	nomeDoItens = append(nomeDoItens, "sabonete")
	nomeDoItens = append(nomeDoItens, "carne")
	nomeDoItens = append(nomeDoItens, "feijao")

	compra := model.NewCompra("Mercadinho", time.Now(), nomeDoItens)
	fmt.Println(compra)

}
