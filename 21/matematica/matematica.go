package matematica

// Sum é uma função genérica que aceita dois parâmetros do mesmo tipo T e retorna a soma deles.
func Sum[T int | float64 | float32](a, b T) T {
	return a + b
}

var A int = 10
