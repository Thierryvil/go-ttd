package main

const prefixHelloPortuguese = "Olá, "

func Hello(nome string) string {
	if nome == "" {
		nome = "Mundo!"
	}
	return prefixHelloPortuguese + nome
}

func main() {
}
