package calculadora

import (
	"fmt"
	"io/ioutil"
)

func menu() {
	var entrada, sair int
	var aux float64

	for sair != 1 {
		fmt.Println("Selecione o tipo de operação que deseja fazer:\n1 para adição\n2 para subtração\n3 para multiplicação\n4 para divisão\n5 para soma continua\n6 para historico\n0 para finalizar")
		fmt.Scanln(&entrada)
		switch entrada {
		case 1:
			aux = adição()
			salvar("\nAdição com resultado:", aux)

		case 2:
			aux = subtração()
			salvar("\nSubtração com resultado:", aux)

		case 3:
			aux = multiplicação()
			salvar("\nMultiplicação com resultado:", aux)

		case 4:
			aux = divisão()
			salvar("\nDivisão com resultado:", aux)

		case 5:
			aux = somacont()
			salvar("\nSoma continua com resultado:", aux)

		case 6:
			imprimir, err := ioutil.ReadFile("historico.txt")
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(imprimir))

		case 0:
			sair = 1
		}
	}
}
