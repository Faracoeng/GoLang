Evitar práticas para ter menos problemas ao utilizar Goroutines

As requisições demoram bastante para ocorrer.

Adicionando apenas `go fetchURL(url)`, a função ja é executada em go routine, porém o codigo encerra antes do final da execução das funções.

Então adiocna um `var wg sync.WaitGroup`


