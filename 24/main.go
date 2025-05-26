package main

// Estudos com condicionais
// Condicionais são estruturas de controle que permitem executar diferentes blocos de código com base em condições.
// No Go, as condicionais são implementadas usando a estrutura if e else.
// Switch é uma alternativa ao if que permite comparar uma variável com vários valores possíveis de forma mais legível.
import ()

func main() {
	a := 10
	b := 20
	c := 30

	if a > b && a > c {
		println("A é o maior número") // Output: A é o maior número
	}
	if b > a && b > c {
		println("B é o maior número") // Output: B é o maior número
	} else if c > a && c > b {
		println("C é o maior número") // Output: C é o maior número
	} else {
		println("Todos são iguais") // Output: Todos são iguais
	}

	if a < b || a < c {
		println("A é menor que B e C") // Output: A é menor que B e C
	} else {
		println("A não é menor que B e C") // Output: A não é menor que B e C
	}

	if a > b {
		println("A é maior que B") // Output: A é maior que B
	} else if a < b {
		println("A é menor que B") // Output: A é menor que B
	} else {
		println("A é igual a B") // Output: A é igual a B
	}


	switch a {
	case 10:
		println("A é igual a 10") // Output: A é igual a 10
	case 20:
		println("A é igual a 20") // Output: A é igual a 20
	default:
		println("A não é igual a 10 nem a 20") // Output: A não é igual a 10 nem a 20
	}
}
