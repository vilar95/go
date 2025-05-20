package main

import (
	"fmt"
)

// Estudos de Structs
type Address struct {
	Street        string
	NumberAddress int
	City          string
	State         string
}
// Structs aninhadas
// A struct Client possui uma struct Address dentro dela
type Client struct {
	Name   string
	Age    int
	Active bool
	Address // A struct Address é aninhada dentro da struct Client
	// Isso significa que os campos da struct Address podem ser acessados diretamente através da struct Client
}

func main() {
	// Criação de um cliente
	eduardo := Client{
		Name:   "Eduardo",
		Age:    29,
		Active: true,
	}
	// Atualizando os dados do cliente
	// Acessando os campos da struct Address através da struct Client
	eduardo.Active = false
	eduardo.Street = "Rua das Flores"
	eduardo.NumberAddress = 123
	eduardo.City = "São Paulo"
	eduardo.State = "SP"
	// Exibe os dados do cliente
	fmt.Printf("\n Nome: %s\n Idade: %d\n Ativo: %t\n Rua: %s\n Número: %d\n Cidade: %s\n Estado: %s\n\n", eduardo.Name, eduardo.Age, eduardo.Active, eduardo.Street, eduardo.NumberAddress, eduardo.City, eduardo.State)
}
