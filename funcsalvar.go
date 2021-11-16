package calculadora

import (
	"fmt"
	"os"
)

func salvar(tipo string, valor float64) {
	arquivo := "historico.txt"
	aux2 := fmt.Sprintf("%f", valor)

	f, err := os.OpenFile(arquivo, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(tipo); err != nil {
		fmt.Println(err)
	}
	if _, err = f.WriteString(aux2); err != nil {
		fmt.Println(err)
	}
}
