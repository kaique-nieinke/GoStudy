package model

import "time"

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDoItens []string) *Compra {
	var itens []ItemDaCompra

	for _, Nome := range nomeDoItens {
		itens = append(itens, ItemDaCompra{Nome})
	}
	return &Compra{
		Mercado: mercado,
		Data:    time.Now(),
		Itens:   itens,
	}
}
