package main

import (
	"fmt"
)

func main() {
	// Cria um slice de inteiros com 10 elementos (de 10 a 100, com incremento de 10)
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// Imprime o tamanho (len), a capacidade (cap) e os valores do slice
	// Saída: len=10 cap=10 [10 20 30 40 50 60 70 80 90 100]
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	// Cria um slice vazio ([:0]) a partir do original, mantendo a mesma capacidade
	// Saída: len=0 cap=10 []
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// Cria um slice com os 4 primeiros elementos (do índice 0 até o 3)
	// Saída: len=4 cap=10 [10 20 30 40]
	fmt.Printf("len=%d cap=%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// Cria um slice a partir do índice 2 até o final
	// A capacidade diminui pois começa do índice 2
	// Saída: len=8 cap=8 [30 40 50 60 70 80 90 100]
	fmt.Printf("len=%d cap=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// Cria um slice a partir do índice 8 até o final (últimos 2 elementos)
	// Saída: len=2 cap=2 [90 100]
	fmt.Printf("len=%d cap=%d %v\n", len(slice[8:]), cap(slice[8:]), slice[8:])

	// Adiciona o valor 110 ao final do slice
	// Como a capacidade seria excedida, o Go cria um novo slice com capacidade maior (dobra para 20)
	slice = append(slice, 110)
	// Imprime o novo tamanho, capacidade e valores
	// Saída: len=11 cap=20 [10 20 30 40 50 60 70 80 90 100 110]
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
