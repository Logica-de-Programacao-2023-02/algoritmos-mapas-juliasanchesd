package main

import "fmt"

func SomarValoresMapa(mapa map[string]int) int {
	soma := 0
	for _, valor := range mapa {
		soma += valor
	}
	return soma
}

func main() {
	valores := map[string]int{"a": 10, "b": 20, "c": 30}

	resultado := SomarValoresMapa(valores)

	fmt.Println("Soma dos valores do mapa:", resultado)
}
