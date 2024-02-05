package main

import (
	"errors"
	"fmt"
)

/*
	: erros encapsulado.
*/

func primeiroErro() error {
	return fmt.Errorf("erro original: algo deu errado na primeira função")
}

func segundoErro() error {
	primeiro := primeiroErro()
	if primeiro != nil {
		//segundo := errors.New("falha na segunda func")
		return fmt.Errorf("falha na segunda func %w", primeiro)
		//return errors.Join(segundo, primeiro)
	}
	return nil
}

func main() {
	err := segundoErro()
	// fmt.Println(err)

	innerError := errors.Unwrap(err)
	fmt.Println(innerError)
}
