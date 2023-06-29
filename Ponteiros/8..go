package main

import "fmt"

// Função para modificar o valor de uma variável usando um ponteiro
func modificarVariavel(p *int) {
	*p = 42
}

func main() {
	// Exemplo de uso da função
	var valor int = 10

	modificarVariavel(&valor)

	fmt.Println(valor) // Output: 42
}
