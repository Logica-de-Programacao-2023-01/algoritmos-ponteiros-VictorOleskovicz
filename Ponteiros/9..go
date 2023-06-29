package main

import "fmt"

// Struct Livro
type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

// Função para aplicar um desconto de 10% no preço do livro
func aplicarDescontoLivro(l *Livro) {
	l.Preco = l.Preco * 0.9
}

func main() {
	// Exemplo de uso da função
	livro := Livro{
		Titulo: "Livro 1",
		Autor:  "Autor 1",
		Preco:  50.0,
	}

	aplicarDescontoLivro(&livro)

	fmt.Println(livro.Preco) // Output: 45.0
}
