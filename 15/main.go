package main

import ()

// Estudos de ponteiros
// Ponteiros são variáveis que armazenam o endereço de outra variável na memória
// Isso permite que você manipule o valor da variável original sem precisar copiá-la
// Ponteiros são úteis quando você deseja passar uma variável para uma função sem copiá-la
// Isso pode economizar memória e melhorar o desempenho do seu programa
// Ponteiros são representados pelo símbolo "&" antes do nome da variável
// O símbolo "*" é usado para desreferenciar o ponteiro, ou seja, acessar o valor armazenado no endereço de memória

func main() {
	// Memoria -> Endereco -> Valor
	a := 10
	println("Valor de a:", a)
	
	// Endereço de "a"
	var ponteiro *int = &a
	println("O espaco de memoria de a:", ponteiro)

	// O valor de a foi alterado para 20
	*ponteiro = 20
	println("Valor de a:", a)

	// O "b" é um ponteiro que aponta para o mesmo endereço de "a"
	// Isso significa que se você alterar o valor de "b", o valor de "a" também será alterado
	b := &a
	println("Valor de b:", b)
	println("O espaco de memoria de b:", &b)
}