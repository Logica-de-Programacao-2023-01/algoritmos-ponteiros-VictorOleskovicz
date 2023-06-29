package main

// Função para atualizar o valor do inteiro para a soma dos n primeiros números naturais
func atualizarSomaNaturais(p *int, n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	*p = sum
}

func main() {
	num := 0
	atualizarSomaNaturais(&num, 5)
}
