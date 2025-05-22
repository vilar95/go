package main

import "fmt"

// Estudos de Ponteiros
// Funcionalidade de um banco
type BankAccount struct {
	balance int
}
// O método deposit é um método de valor, o que significa que ele não altera o valor original do saldo.
func (bankAccount BankAccount) deposit(value int) int {
	bankAccount.balance += value
	fmt.Printf("O saldo da conta é: %v reais, com o depósito (Dentro da função)\n", bankAccount.balance)
	return bankAccount.balance
}
// O método loanSimulation é um método de ponteiro, o que significa que ele altera o valor original do saldo.
// Isso é útil quando você deseja modificar o valor do saldo da conta original dentro da função.
func (bankAccount *BankAccount) loanSimulation(value int) int {
	bankAccount.balance += value
	fmt.Printf("O saldo da conta é: %v reais, com a solicitação de empréstimo (Dentro da função)\n", bankAccount.balance)
	return bankAccount.balance
}

func main() {
	account := BankAccount{
		balance	: 100,
	}
	account.loanSimulation(200)
	fmt.Printf("O saldo da conta é: %v reais, com a solicitação de empréstimo\n", account.balance)

	account.deposit(500)
	fmt.Printf("O saldo da conta é: %v reais, com o depósito\n", account.balance)
}
