package main

import(
	"fmt"
	"time"
)

func incrementar(contador *int){
	// Aqui eu estou incrementando o valor do contador
	// No caso de um tempo muito demorado, varias goRoutines podem mecher na mesma variavel atraves de seu ponteiro dando valores equivocados
	for i:= 0; i<1000;i++{
		*contador++
	}
}

func main(){
	contador := 0

	for i := 0; i<10;i++{
		// Não passar ponteiro como parametro em uma Goroutine
		incrementar(&contador)
	}
	fmt.Println("final:", contador)
	time.Sleep(1 * time.Second)
}