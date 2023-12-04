package main

import (
	"fmt"
)

func FrequenciaCaracteres(s string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range s {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	texto := "hello world"

	resultado := FrequenciaCaracteres(texto)

	fmt.Println("Frequência de caracteres:")
	for char, freq := range resultado {
		fmt.Printf("%c: %d\n", char, freq)
	}
}
