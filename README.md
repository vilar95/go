# Estudos em GoLang 🚀

Este repositório reúne exemplos práticos e anotações dos principais conceitos da linguagem Go, organizados por temas e pastas numeradas. Cada pasta contém um arquivo `main.go` com exemplos comentados. Abaixo, um resumo dos tópicos estudados:

## 1. Hello World e Introdução 👋

- Exemplo clássico de Hello World com arte ASCII.

## 2. Tipos Básicos e Variáveis 🔤

- Declaração de constantes, variáveis globais e locais.
- Tipos primitivos: `int`, `bool`, `string`, `float64`.

## 3. Type Alias 🏷️

- Criação de tipos personalizados usando `type`.

## 4. Formatação de Saída 🖨️

- Uso de `fmt.Printf` para formatar valores, tipos e casas decimais.

## 5. Arrays 📦

- Declaração, inicialização e iteração sobre arrays.

## 6. Slices ✂️

- Criação, manipulação, capacidade e tamanho de slices.
- Uso de `append` para adicionar elementos.

## 7. Mapas 🗺️

- Criação, manipulação, deleção e iteração sobre mapas.

## 8. Funções com Retorno de Erro ⚠️

- Funções que retornam múltiplos valores, incluindo erros.
- Uso de `errors.New` e tratamento de erros.

## 9. Funções Variádicas ➕

- Funções que aceitam número variável de argumentos (`...int`).

## 10. Funções Anônimas 🙈

- Definição e uso de funções anônimas atribuídas a variáveis.

## 11-14. Structs e Interfaces 🏗️

- Definição de structs simples e aninhadas.
- Métodos em structs.
- Implementação de interfaces.
- Métodos com ponteiros para modificar structs.

## 15-17a. Ponteiros 📍

- Conceito de ponteiros, endereçamento e desreferenciação.
- Métodos de valor vs. métodos de ponteiro.
- Exemplo prático: simulação de conta bancária.

## 18-19. Interface Vazia e Type Assertion 🎭

- Uso de `interface{}` para aceitar qualquer tipo.
- Type assertion para converter interface em tipos específicos.

## 20-21. Generics e Pacotes 🧩

- Uso de generics para funções e tipos genéricos.
- Criação e importação de pacotes próprios (`matematica`).

## 22. Módulos e Pacotes Externos 📦🌐

- Uso de módulos (`go.mod`).
- Importação e uso do pacote externo `github.com/google/uuid` para geração de UUIDs.

## 23. Loops e Controle de Fluxo 🔁

- Estruturas de repetição: `for`, `range`, `while` (simulado).
- Uso de `break` para sair de loops.

## 24. Condicionais 🔀

- Estruturas `if`, `else if`, `else` e `switch`.

---

## Observações 📝

- Alguns exemplos requerem Go 1.18+ para suporte a generics.
- Para exemplos com pacotes externos, execute `go mod tidy` na pasta para baixar as dependências.

---

Estes exemplos servem como referência rápida para conceitos fundamentais de Go. Sinta-se à vontade para explorar, modificar e experimentar! 🎉
