package operacoes

import (
	"fmt"
)

func adição() float64 {
	var a, b, resultado int
	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo numero")
	fmt.Scanln(&b)
	resultado = a + b
	fmt.Println("O resultado é:", resultado)
	return float64(resultado)
}

func subtração() float64 {
	var a, b, resultado int
	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo numero")
	fmt.Scanln(&b)
	resultado = a - b
	fmt.Println("O resultado é:", resultado)
	return float64(resultado)

}

func multiplicação() float64 {
	var a, b, resultado int
	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo numero")
	fmt.Scanln(&b)
	resultado = a * b
	fmt.Println("O resultado é:", resultado)
	return float64(resultado)

}

func divisão() float64 {
	var a, b, resultado float64
	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo numero")
	fmt.Scanln(&b)
	if b == 0 {
		println("Não é possivel dividir por zero")
	} else {
		resultado = a / b
		fmt.Println("O resultado é:", resultado)
	}
	return resultado

}

func somacont() float64 {
	var aux, resultadocont int
	for {
		println("\nDigite o numero a ser somado, aperte 0 para finalizar")
		fmt.Scanln(&aux)
		if aux == 0 {
			fmt.Println("Resultado final:", resultadocont)
			break
		} else {
			resultadocont += aux
			fmt.Println("Resultado parcial:", resultadocont)
		}
	}
	return float64(resultadocont)

}
