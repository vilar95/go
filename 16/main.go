package main

import ()

// Estudos de Ponteiros
// Quando deve utilizar ponteiros?

// Quando você quer passar o endereço de uma variável para uma função.
// Isso é útil quando você deseja modificar o valor da variável original dentro da função.
func sum(a, b int) int {
	// A variável "a" é inicializada com o valor 50
	// Isso significa que o valor de "a" dentro da função não afetará o valor de "a" fora da função,
	// pois "a" é passado por valor (cópia).
	a = 50
	return a + b
}

func subtract(a, b *int) int {
	// O valor apontado por "a" é alterado para 70.
	// Como "a" é um ponteiro, essa alteração afeta o valor original fora da função.
	*a = 70
	return *a - *b
}

func main() {
	a := 10
	b := 20

	// O valor de "a" dentro da função será alterado para 50,
	// mas fora da função "a" permanece 10.
	println("Valor da soma de:", sum(a, b))
	// O valor de "a" fora da função ainda será 10.
	println("Valor de a:", a)
	// O valor apontado por "a" será alterado para 70 dentro da função,
	// e essa alteração afeta "a" fora da função.
	println("Valor da subtração:", subtract(&a, &b))
}
