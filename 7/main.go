package main

import (
	"fmt"
)
// O código abaixo é um exemplo de como criar um mapa em Go
func main() {
	// Criando um mapa com chaves do tipo string e valores do tipo int
	salarios := map[string]int{
		"Wesley": 1000,
		"João":   2000,
		"Maria":  3000,
	}
	// Imprimindo o mapa
	fmt.Printf("O mapa salarios é: %v\n", salarios)
	// Deletando o elemento "Wesley" do mapa
	delete(salarios, "Wesley")
	salarios["Eduardo"] = 10000
	fmt.Println("O salario de Wesley foi deletado")
	// Imprimindo o mapa sem o elemento "Wesley"
	fmt.Printf("O mapa salarios é: %v\n", salarios)
	// Verificando o slario de Wesley
	fmt.Printf("O salario de Wesley é %d\n", salarios["Wesley"])

	// Imprimindo o salario de cada pessoa no mapa
	// O range percorre o mapa e retorna a chave e o valor
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

	// salarios2 := make(map[string]int)
	// salarios2["Pedro"] = 1000
	// salarios2["Lucas"] = 2000
	// salarios2["Paula"] = 3000

	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("O mapa salarios2 é: %v\n", salarios2)
	// }
	
}
