package main

import (
	"errors"
	"fmt"
)

func main() {
	// Exemplo de variáveis
	valor, err := sum(10, 20)
	// Se a soma for maior que 50, retorna um erro
	// Se não, retorna o valor da soma
	fmt.Println(valor)
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
// função que soma dois números inteiros
// e retorna uma string com o resultado e um erro
func sum(a, b int) (string, error) {
	calculation := a + b
	// Cria uma string com o resultado da soma
	// %d é um placeholder para um número inteiro
	msg := fmt.Sprintf("A soma dos valores é %d", calculation)

	// Se a soma for maior que 50, retorna um erro
	if calculation >= 50 {
		return msg, errors.New("A soma deve ser menor que 50")
	}
	// Se não, retorna o valor da soma e o erro como nulo
	return msg, nil
}
