package main

import (
	"fmt"
	"projeto2/banco/clientes"
	"projeto2/banco/contas"
)

func PagarBoleto(conta VerificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoBruno := contas.ContaCorrente{
		Titular: clientes.Cliente{
			Nome:      "Bruno",
			CPF:       "12345678900",
			Profissao: "Desenvolvedor",
		},
		Agencia: 123,
		Conta:   123456,
	}

	contaDoBruno.Depositar(100)
	PagarBoleto(&contaDoBruno, 60)

	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 100)

	fmt.Println(contaDoBruno.VisualizarSaldo())
	fmt.Println(contaDaLuisa.VisualizarSaldo())
}
