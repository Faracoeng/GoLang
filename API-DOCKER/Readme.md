## Desenvolvimentode API estruturada com padrão de projeto e dockerizada

Iniciar projeto com:

> go mod init go-api


Instalar o framework usado para desenvolver API:


> go get github.com/gin-gonic/gin


Subir anco de dados postgres pelo `docker-compose.yml`, a tabela deve ser:

```sql
CREATE TABLE IF NOT EXISTS products (
    id_product SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
```

Criar na raiz do repositório `model`, para armazenar os models da API.

Estrutura do produto:

```go
package model

type Product struct {
	ID		  int     `json:"id_product"`
	Name		  string  `json:"name"`
	Price		  float64 `json:"price"`
}
```

Criar na raiz do repositório `controller`, receber as requisições da API, e dar os retornos.

Criar na raiz do repositório `use-case`, para as regras e negócio da API.
