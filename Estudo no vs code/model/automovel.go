package main

type Automovel struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Automovel   Automovel
	Cilindradas int
}

type Carro struct {
	Automovel
	QuantidadeDePortas int
	Potencia           int
	ArCondicionado     bool
}
