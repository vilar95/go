package main

// Estudos com looping e controle de fluxo
// Looping e controle de fluxo são conceitos fundamentais em programação que permitem executar um bloco de código várias vezes ou tomar decisões com base em condições.
// No Go só existem dois tipos de loops: for e range.
import (

)
func main() {
	// Exemplo de loop for
	for i := 0; i <= 10; i++ {
		println("Loop for:", i) // Output: Loop for: 0, Loop for: 1, ..., Loop for: 10
	}

	// Exemplo de loop range
	numbers := []string{"Um", "Dois", "Três", "Quatro", "Cinco"}
	for index, value := range numbers {
		println("Index:", index, "Value:", value) // Output: Index: 0 Value: Um, Index: 1 Value: Dois, ..., Index: 4 Value: Cinco
	}

	i := 0
	for i < 5 {
		println("Loop while:", i) // Output: Loop while: 0, Loop while: 1, ..., Loop while: 4
		i++
	}

	for {
		println("Loop infinito") // Cuidado! Isso criará um loop infinito, interrompa manualmente
		break // Use break para sair do loop infinito
	}
}