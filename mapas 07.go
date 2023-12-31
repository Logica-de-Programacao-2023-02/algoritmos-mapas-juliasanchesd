package main

import (
	"fmt"
	"strings"
)

func ContagemLetrasPorPalavra(frase string) map[string]map[rune]int {
	palavras := strings.Fields(frase)
	resultado := make(map[string]map[rune]int)

	for _, palavra := range palavras {
		resultado[palavra] = contagemLetras(palavra)
	}

	return resultado
}

func contagemLetras(palavra string) map[rune]int {
	contagem := make(map[rune]int)

	for _, char := range palavra {
		contagem[char]++
	}

	return contagem
}

func main() {
	frase := "hello world"
	resultado := ContagemLetrasPorPalavra(frase)

	fmt.Println("Contagem de letras por palavra:")
	for palavra, contagem := range resultado {
		fmt.Printf("%s: %v\n", palavra, contagem)
	}
}
