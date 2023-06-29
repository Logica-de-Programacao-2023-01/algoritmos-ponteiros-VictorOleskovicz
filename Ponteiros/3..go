package main

// Função para inverter o valor da string
func inverterString(p *string) {
	runes := []rune(*p)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*p = string(runes)
}

func main() {
	str := "Hello, World!"
	inverterString(&str)
}
