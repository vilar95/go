package main

import (
	"fmt"
)

func main() {
	// Exemplo de variáveis
	// Chama a função sum com um número variável de argumentos
	fmt.Println(sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500))
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
