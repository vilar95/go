package main

// Estudos com pacotes e módulos
// Pacotes são uma maneira de organizar o código em Go, permitindo que você divida seu programa em partes menores e reutilizáveis.
// Módulos são coleções de pacotes que podem ser versionados e compartilhados.

import (
	"GO/matematica" // Importa o pacote matematica que contém funções genéricas
	"fmt"
)

func main() {
	result := matematica.Sum(10, 20) // Chama a função sum do pacote math com dois inteiros
	fmt.Println("O resultado da soma é:", result) // Output: 3

	fmt.Println("O valor de A é:", matematica.A) // Output: 10
}