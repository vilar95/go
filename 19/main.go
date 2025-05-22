package main

import "fmt"

// Estudos com type assertion
// A type assertion é uma maneira de verificar o tipo de um valor armazenado em uma interface vazia.
// Isso é útil quando você deseja realizar operações específicas para esse tipo de dado.

func main() {
	
	var myVar interface{} = "Eduardo Vilar"
	
	// (0x4f19a0,0x51beb0) Quando você imprime uma variável do tipo interface{}, o Go imprime o endereço de memória onde a variável está armazenada.
	// Isso é útil para depuração, mas não é muito legível.
	println("myVar é do tipo:", myVar)

	// É precico fazer uma type assertion para converter a interface vazia em um tipo específico.
	// Isso significa que você está dizendo ao Go: "Eu sei que essa interface contém um valor do tipo string, então me dê esse valor como uma string."
	println("myVar é:", myVar.(string))

	res, ok := myVar.(int)
	// O valor de res é 0 e o valor de ok é false
	fmt.Printf("O valor de res é: %v e o valor de ok é: %v\n", res, ok)

	// panic: interface conversion: interface {} is string, not int
	// Isso significa que você está tentando converter uma string em um inteiro, o que não é permitido.
	res2 := myVar.(int)
	fmt.Printf("O valor de res2 é: %v\n", res2)
}