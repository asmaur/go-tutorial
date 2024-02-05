package main

import (
	"fmt"
)

/*
	error: lidar com erros de forma diferente
	-> sem try-catch ( nem except)
	-> as função retornam o erro como ultimo valor
	-> a mensagem de erro é sempre em minuscula e sem ponto final
*/

type customError struct {
	code    int
	message string
}

func (c customError) Error() string {
	return fmt.Sprintf("Erro %d: %s\n", c.code, c.message)
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// return 0, errors.New("divisão por zero")
		return 0, customError{code: 12, message: "divisão por zero"}
	}
	return a / b, nil
}

func main() {
	v, err := dividir(10, 0)

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	fmt.Println("Resultado: ", v)
}
