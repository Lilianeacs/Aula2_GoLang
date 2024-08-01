package contas

import "projeto2/banco/clientes"

type ContaCorrente struct {
	Titular clientes.Cliente
	Agencia int
	Conta   int
	saldo   float64
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= conta.saldo && valorDoSaque > 0 {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (conta *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		conta.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", conta.saldo
	} else {
		return "O valor do depósito é menor ou igual a zero!", conta.saldo
	}
}

func (conta *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < conta.saldo && valorDaTransferencia > 0 {
		conta.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	return false
}

func (conta *ContaCorrente) VisualizarSaldo() float64 {
	return conta.saldo
}
