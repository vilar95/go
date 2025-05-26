package main
// Estudos com pacotes e módulos
// Pacotes são uma maneira de organizar o código em Go, permitindo que você divida seu programa em partes menores e reutilizáveis.
// Módulos são coleções de pacotes que podem ser versionados e compartilhados.
import (
	"fmt"
	"github.com/google/uuid"
)
func main() {
	id := uuid.New()
	fmt.Println("UUID gerado:", id.String()) // Gera e exibe um UUID único
	fmt.Println("UUID versão:", id.Version()) // Exibe a versão do UUID
	fmt.Println("UUID variante:", id.Variant()) // Exibe a variante do UUID
	fmt.Println("UUID bytes:", id[:]) // Exibe os bytes do UUID
	fmt.Println("UUID string:", id.String()) // Exibe o UUID como string
}