package main

import (
	"fmt"
)

// Função para somar dois números
func sum(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	// Solicita os números ao usuário
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)

	// Calcula a soma
	result := sum(num1, num2)

	// Exibe o resultado
	fmt.Printf("A soma de %d e %d é %d\n", num1, num2, result)
}
