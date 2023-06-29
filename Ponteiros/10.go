package main

import "fmt"

// Função para verificar se um número é primo
func isPrimo(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Função para preencher uma slice com os n primeiros números primos
func preencherPrimos(s *[]int, n int) {
	count :=
