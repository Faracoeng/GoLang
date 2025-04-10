package main
import (
	"fmt"
	"net/http"
	"sync"
)
func fetchURL(wg *sync.WaitGroup, url string) {
    defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao requisitar %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Requisição para %s concluida com status %s\n", url, resp.Status)
}
func oldmain(){
	var wg sync.WaitGroup

	urls := []string{
		"https://golang.com",
		"https://github.com",
		"https://stackoverflow.com",
	}
	for _, url := range urls {
		// A cada iteração do for, uma goroutine é disparada, e registrada no WaitGroup
        wg.Add(1)
		go fetchURL(&wg, url)
	}
	// Aqui eu digo parao Go nao encerrar ate todas as GoRoutines do waitGroup estiverem sido encerradas
	wg.Wait()
	fmt.Println("Requisições concluidas")
}