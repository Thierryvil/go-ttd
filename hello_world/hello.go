package main

const prefixHelloPortuguese = "Olá, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFrench = "Bonjour, "

const spanish = "spanish"
const french = "french"

func Hello(name, language string) string {
	if name == "" {
		name = "Mundo!"
	}

	return prefixHello(language) + name

}

func prefixHello(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixHelloFrench
	default:
		prefix = prefixHelloPortuguese
	}
	return
}
