package contas

import "projeto2/banco/clientes"

type ContaPoupanca struct {
	Titular                  clientes.Cliente
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (conta *ContaPoupanca) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= conta.saldo && valorDoSaque > 0 {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (conta *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		conta.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", conta.saldo
	} else {
		return "O valor do depósito é menor ou igual a zero!", conta.saldo
	}
}

// func (conta *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
// 	if valorDaTransferencia < conta.saldo && valorDaTransferencia > 0 {
// 		conta.saldo -= valorDaTransferencia
// 		contaDestino.Depositar(valorDaTransferencia)
// 		return true
// 	}
// 	return false
// }

func (conta *ContaPoupanca) VisualizarSaldo() float64 {
	return conta.saldo
}
