package main

import (
	"banco/contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(float64) string
}

func main() {
	contaDaMaria := contas.ContaPoupanca{}
	contaDaMaria.Depositar(200)
	fmt.Println(contaDaMaria)

	pagarBoleto(&contaDaMaria, 100)
	fmt.Println(contaDaMaria)

	// contaDoJoao := Conta{"João", 123456789, 1234, 1025.88}
	// contaDoJoao2 := Conta{"João", 123456789, 1234, 1025.88}

	// fmt.Println("Deve ser verdadeiro")
	// fmt.Println(contaDoJoao == contaDoJoao2)

	// var contaDaMaria *Conta
	// contaDaMaria = new(Conta)
	// contaDaMaria.titular = "maria"
	// contaDaMaria.saldo = 5000

	// var contaDaMaria2 *Conta
	// contaDaMaria2 = new(Conta)
	// contaDaMaria2.titular = "maria"
	// contaDaMaria2.saldo = 5000

	// fmt.Println("Deve ser Falso, os * são referencias aos ponteiros da memória. Portanto, o ponteiro da memória da contaMaria e Conta Maria2 são diferentes. ")
	// fmt.Println(contaDaMaria == contaDaMaria2)

	// fmt.Println("Deve ser verdadeiro, porque agora se compara os valores ")
	// fmt.Println(*contaDaMaria == *contaDaMaria2)

}
