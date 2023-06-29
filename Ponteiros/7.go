package main

import "fmt"

// Struct Conta
type Conta struct {
	Saldo   float64
	Titular string
}

// Função para adicionar um valor ao saldo da conta
func adicionarSaldo(c *Conta, valor float64) {
	c.Saldo += valor
}

func main() {
	// Exemplo de uso da função
	conta := Conta{
		Saldo:   100.0,
		Titular: "João",
	}

	adicionarSaldo(&conta, 50.0)

	fmt.Println(conta.Saldo) // Output: 150.0
}
