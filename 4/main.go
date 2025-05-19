package main

import (
	"fmt"
)

var (
	pi float64 = 3.1415926
)

func main() {
 fmt.Printf("O tipo de PI é %T", pi)
 fmt.Printf("\nO valor de PI é %v", pi)
 fmt.Printf("\nO valor de PI é %.2f", pi)
 fmt.Printf("\nO valor de PI é %.4f", pi)
 fmt.Printf("\nO valor de PI é %.6f", pi)
 fmt.Printf("\nO valor de PI é %.8f", pi)
}
