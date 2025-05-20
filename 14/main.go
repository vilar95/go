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
// O struct Client implementa a interface People.
// Isso significa que o struct Client deve ter todos os métodos da interface People.
type People interface {
	// As interfaces só podem ter métodos
	// As interfaces não podem ter Atributos e Funções
	disableClient()
}

// O método disableClient é um método da struct Client.
// Ele altera o valor do campo Active para false.
// O ponteiro "*" é necessário para alterar o valor do campo Active.
// O método é chamado através de um ponteiro para a struct Client.
func (c *Client) disableClient (){
	// Desabilita o cliente
	c.Active = false
	fmt.Printf("O cliente %s foi desativado com sucesso!", c.Name)
}

func main() {
	// Cria um cliente
	eduardo := Client{
		Name:   "Eduardo",
		Age:    29,
		Active: true,
	}
	// Desabilita o cliente
	eduardo.disableClient()
	// Altera os dados do cliente
	eduardo.Street = "Rua das Flores"
	eduardo.NumberAddress = 123
	eduardo.City = "São Paulo"
	eduardo.State = "SP"
	// Exibe os dados do cliente
	fmt.Printf("\n Nome: %s\n Idade: %d\n Ativo: %t\n Rua: %s\n Número: %d\n Cidade: %s\n Estado: %s\n\n", eduardo.Name, eduardo.Age, eduardo.Active, eduardo.Street, eduardo.NumberAddress, eduardo.City, eduardo.State)
}
