package main

import "fmt"

// Estudos com interfaces vazias

func main() {
	// A interface vazia é um tipo que pode armazenar qualquer valor.
	// Isso significa que você pode armazenar qualquer tipo de dado dentro dela.
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	
	fmt.Println("Print sem condições para verificar o tipo do valor armazenado na interface vazia")
	showType(x)
	showType(y)

	fmt.Println("\nPrint com condições para verificar o tipo do valor armazenado na interface vazia")
	showTypeInt(x)
	showTypeInt(y)
}
// A função showType recebe um parâmetro do tipo interface{}.
// Isso significa que ela pode receber qualquer tipo de dado.	
func showType(i interface{}) {
	fmt.Printf("O tipo é: %T e o valor é %v\n", i, i)
}

// Dica: Adicione condições para verificar o tipo do valor armazenado na interface vazia. Para isso não de ruim manito
// A função showTypeInt verifica se o valor armazenado na interface é do tipo int.
// Isso é útil quando você deseja realizar operações específicas para esse tipo de dado.
func showTypeInt(i interface{}) {
	if value, ok := i.(int); ok {
		fmt.Printf("O tipo é: %T e o valor é %v\n", value, value)
	} else {
		fmt.Println("O valor não é um inteiro.")
	}
}