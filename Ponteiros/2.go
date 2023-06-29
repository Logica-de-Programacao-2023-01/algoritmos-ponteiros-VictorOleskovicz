package main

// Função para verificar se o inteiro é par ou ímpar e atualizar o valor para 0 (par) ou 1 (ímpar)
func verificarParImpar(p *int) {
	if *p%2 == 0 {
		*p = 0
	} else {
		*p = 1
	}
}

func main() {
	num := 7
	verificarParImpar(&num)
}
