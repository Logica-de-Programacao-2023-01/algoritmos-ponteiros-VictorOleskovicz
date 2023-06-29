package main

import "fmt"

// Struct Livro
type Livro struct {
	Titulo string
	Autor  string
}

// Função para atualizar o título do livro para "Desconhecido" se o autor for "Anônimo"
func atualizarTituloLivro(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Titulo = "Desconhecido"
	}
}

func main() {
	// Exemplo de uso da função
	livro := Livro{
		Titulo: "Livro 1",
		Autor:  "Anônimo",
	}

	atualizarTituloLivro(&livro)

	fmt.Println(livro.Titulo) // Output: Desconhecido
}
