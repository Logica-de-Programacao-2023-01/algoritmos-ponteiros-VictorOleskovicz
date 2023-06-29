package main

import "math"

// Função para atualizar o valor da variável float64 para a média aritmética entre seu valor atual e Pi
func atualizarMediaPi(p *float64) {
	*p = (*p + math.Pi) / 2
}

func main() {
	num := 3.0
	atualizarMediaPi(&num)
}
