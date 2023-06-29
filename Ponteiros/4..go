package main

import "strconv"

// Função para atualizar o valor da variável para a soma dos dois últimos dígitos
func atualizarSomaUltimosDigitos(p *int) {
	digits := strconv.Itoa(*p)
	length := len(digits)
	if length >= 2 {
		lastDigit, _ := strconv.Atoi(string(digits[length-1]))
		secondLastDigit, _ := strconv.Atoi(string(digits[length-2]))
		*p = lastDigit + secondLastDigit
	}
}

func main() {
	num := 1234
	atualizarSomaUltimosDigitos(&num)
}
