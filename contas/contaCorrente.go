package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular          clientes.Pessoa
	NumeroDaConta    int
	NnumeroDaAgencia int
	saldo            float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && c.saldo >= valorDoSaque {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saque recusado"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito com sucesso", c.saldo
	} else {
		return "Depósito recusado", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) bool {
	if c.saldo >= valorDaTransferencia && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return true
	} else {
		return false
	}
}

func (c ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
