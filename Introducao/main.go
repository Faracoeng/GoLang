package main

import (
	"calculadoraTeste/calculadora"
	"fmt"
)

func main(){
	//var a int= 5
	a := 7
	b := 0

	resultado, err := calculadora.Dividir(a, b)	

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Resultado: ", resultado)
}



