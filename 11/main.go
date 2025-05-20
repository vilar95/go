package main

import "fmt"

// Estudos de Structs
type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	
	eduardo := Client{
		Name:   "Eduardo",
		Age:    29,
		Active: true,
	}
	eduardo.Active = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", eduardo.Name, eduardo.Age, eduardo.Active)
}
