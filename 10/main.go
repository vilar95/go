package main

import (
	"fmt"
)

func main() {
	// A função anônima é atribuída a uma variável chamada total
	total := func ()  int {
		return sum(10, 20, 30, 40, 50, 60) *2
	}()
	//
	fmt.Println(total)
}
// função que soma números inteiros
func sum(numeros ...int) int {
	// Cria uma variável total com valor 0
	total := 0
	// Itera sobre os números e soma cada um deles ao total
	// O operador ... indica que a função pode receber um número variável de argumentos
	for _, numero := range numeros {
		total += numero
	}
	return total
}
