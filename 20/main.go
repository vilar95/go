package main

import "fmt"

// Tipos genericos
// Generics são uma maneira de escrever código que pode funcionar com diferentes tipos de dados.
// Isso é útil quando você deseja criar funções ou estruturas de dados que podem trabalhar com diferentes tipos sem precisar duplicar o código.
// Em Go, você pode usar generics para criar funções e tipos que podem aceitar diferentes tipos de dados.
// Isso é feito usando a palavra-chave "type" e "interface{}".

// Exemplo de tipagem de variável
type myNumber float64

// Exemplo de função genérica que aceita um mapa e retorna a soma dos valores
type number interface {
	~int | ~float64 | ~float32
}
// A função sum recebe um mapa com chaves do tipo string e valores do tipo T, onde T é um tipo que implementa a interface number.
// A interface number é definida para aceitar tipos numéricos como int, float64 e float32.
func sum[T number](m map[string]T) T {
	result := T(0) // Inicializa o resultado com o tipo T, que pode ser int, float64 ou float32
	for _, v := range m {
		result += v
	}
	return result
}

func compare[T number](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// Exemplo de uso da função sum com diferentes tipos de mapas
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := sum(m)
	println("Resultado:", result) // Output: Sum: 6

	m2 := map[string]float64{"x": 10.2, "y": 20.3, "z": 30.4}
	result2 := sum(m2)
	fmt.Printf("Resultado: %.1f\n", result2) // Output: Sum: 60.9

	m3 := map[string]myNumber{"p": 1.5, "q": 2.5, "r": 3.5}
	result3 := sum(m3)
	fmt.Printf("Resultado: %.1f\n", result3) // Output: Sum: 7.5

	println("Comparando 1.0 e 1:", compare(1.0, 1)) // Output: true
	println("Comparando 1.0 e 2:", compare(1.0, 2)) // Output: false
	
}
