package main

import "fmt"

// Estudos de Ponteiros
// Funcionalidade de um banco
type BankAccount struct {
	balance int
}
// O saldo da conta é 0 reais, com o depósito
// *BankAccount é um ponteiro para BankAccount.
// Isso significa que a função pode modificar o valor do saldo da conta original.
// Isso é útil quando você deseja modificar o valor da variável original dentro da função.
func newAccount() *BankAccount {
	return &BankAccount{
		balance: 0,
	}
}
func (bankAccount BankAccount) deposit(value int) int {
	bankAccount.balance += value
	// O saldo da conta é 0 reais, com o depósito (Dentro da função)
	fmt.Printf("O saldo da conta é: %v reais, com o depósito (Dentro da função)\n", bankAccount.balance)
	return bankAccount.balance
}
// O saldo da conta é 0 reais, com a solicitação de empréstimo (Dentro da função)
func (bankAccount *BankAccount) loanSimulation(value int) int {
	bankAccount.balance += value
	fmt.Printf("O saldo da conta é: %v reais, com a solicitação de empréstimo (Dentro da função)\n", bankAccount.balance)
	return bankAccount.balance
}

func main() {
	// Cria uma nova conta bancária
	// O saldo da conta é 0 reais, com o depósito
	account := newAccount()
	// A simulação de empréstimo é feita com o valor de 200 reais
	account.loanSimulation(200)
	fmt.Printf("O saldo da conta é: %v reais, com a solicitação de empréstimo\n", account.balance)
	// O depósito é feito com o valor de 500 reais
	account.deposit(500)
	fmt.Printf("O saldo da conta é: %v reais, com o depósito\n", account.balance)
	// O saldo da conta é 700 reais
}
