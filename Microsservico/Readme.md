## Microsserviço

Dever possuir um servidor HTTP, que através de uma API, deve ser possivel:

- Criar uma nova categoria;
- Listar todas as categorias;
- Obter uma categoria especifica;
- Deletar uma categoria;
- Atualziar uma categoria existente;
- Notificar um broker Kafka que uma categoria foi atualizada. 


### Processo de desenvolvimento

Iniciar com inicialização do mod, para meu caso:

> go mod init `github.com/Faracoeng/GoLang/ms-categories`

Criar o diretório `cmd`, seguindo padrão da comunidade, nela serão adicionados os executáveis. neste caso, o arquivo com a função **main**.

O framework utilizado será o [gin-gonic](https://gin-gonic.com/), com o seguinte exemplo basico de funcionamento:


``` go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

Para instalar dependências, executar:

> go mod tidy

Ao executar o comando abaixo, o servidor HTTP ja deve estar de pé:

> go run cmd/api/main.go

Ao acessar http://localhost:8080/ping deve receber como retorno o pong.

Agora criar na raiz o arquivo `api.http` para obter os endpoints criados.

Para nao precisar fazer um `go run` novamente toda vez que realizar uma alteração no código, pode ser instalado o [Air - Live reload for Go apps](https://github.com/air-verse/air)


Ou direto pelo Go:

> go install github.com/air-verse/air@latest


Criar o diretório `Internal`, para armazenar apenas arquivos que fazem sentido apenas para meu projeto, ou seja recursos internos. Pacotes reutilizáveis para outras aplicações devem estar no diretório `pkg`. O nome do pacote deve ser idêntico ao nome da pasta, e nomes sempre no singular. Ja com o `pkg`, deve armazenar código comum ao projeto, ou seja, conexão com banco de dados, ou classes de objetos que são utilizados por vários projetos.


Criar então dentro de **Internal** o diretório `entities`, e neste diretório criar `category.go`, com o seguinte conteúdo.

``` go
package entities

type Category struct {
	ID   uint    `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```
Uma boa prática em microsseviços, é adicionar regras de negócio nas próprias entidades. Neste caso, nenhum pacote externo deve instanciar a entidade **Category**, então deve possuir a prória função de inicialização.

```go
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
func NewCategory(id int, name string) (*Category, error) {
	category :=  &Category {
		Name: name,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	// Aqui então devem ser adicionados os métodos para manipular a categoria, como adicionar, editar e remover.
	// e definir as regras de negócio
	return category, nil
}
```

Na sequencia criar um diretório para os casos de uso, em **internal**, para criar os scripts em `go` para cada operação/regra de negocio que a API deve ofertar.

Criar no pacote main em `cmd/api` o arquivo para definir as rotas **routes.go**

As requisições para testar a API estão em `api.http`.

Criar o diretório chamado **repositories** que são responsáveis por fornecer métodos para recuperar e armazenar agregados. Eles atuam como uma abstração da camada de persistência. Nele criaremos uma interface para definir os metodos de conexão. A grande vantagem é que utilizando os metodos através da interfacem o script em GO nao sabe se esta persistindo em MySQL, PostgreSQL, MongoDB, torando a aplicação mais escalável, pois caso o DSN seja alterado, nem os controllers, e nem os use cases precisam ser ajustados.


Basta criar então o script de abstração com o banco de dados, o fluxo deve ficar assim:

```
[ HTTP Controller ]
        │
        ▼
[ Use Case (Application) ]
        │
        ▼
[ ICategoryRepository interface ]
        │
        ▼
[ CategoryRepository (Infra, GORM) ]
        │
        ▼
[ MySQL ]
```



### Boas práticas em microsserviços

- Em microservices, é sempre bom manter a regra de negócio definida na entidade e não fazer com um model de banco de dados por exemplo.

- Neste microsserviço, em internal/entities/category.go, existe a struct que representa uma `Categoria`, e para construir uma Categoria, é necessário utilizar uma própria função de category.go chamada func NewCategory(name string) *Category, e nela, definir as regras de negócio, para validar a estrutura criada.
Internal

- Deve armazenar scripts que são utilizados apenas pelo microsserviço, que nao são compartilhados com outros serviços, esses compartilhados deveriam estar em um pacote chamado pkg.

- Entities: Local onde são criadas as estuturas de dados utilizadas pelo microsserviço.

- Services: Local onde são implementadas as regras de negócio do microsserviço.

- É possivel integrar através desta API, notificar para um broker/MENSSAGERIA quando uma categoria é criada/atualizada.

Referência: [Youtube](<https://www.youtube.com/watch?v=DVZ3hs3Bq34)>)


### Domain (Domínio)
O termo **domain** refere-se ao domínio do problema, ou seja, à lógica de negócios e às regras que são centrais para o problema que o software está resolvendo. Em DDD, o domínio é dividido em várias partes, incluindo entidades, agregados, repositórios e serviços de domínio.

- Entidades (Entities): São objetos que têm identidade própria e que permanecem constantes ao longo do tempo. Por exemplo, um "Cliente" ou "Pedido" são entidades que têm propriedades e métodos próprios.
Objetos de Valor (Value Objects): São objetos que não têm identidade própria e são definidos apenas por seus atributos. Por exemplo, uma "Data" ou "Endereço".
- Agregados (Aggregates): São grupos de entidades e objetos de valor que são tratados como uma única unidade de consistência. Por exemplo, um "Pedido" pode ser um agregado que contém várias "Linhas de Pedido".
- Repositórios (Repositories): São responsáveis por fornecer métodos para recuperar e armazenar agregados. Eles atuam como uma abstração da camada de persistência.
Serviços de Domínio (Domain Services): São operações que não pertencem a nenhuma entidade ou objeto de valor específico, mas que fazem parte da lógica de negócios. Eles encapsulam regras de negócios que envolvem várias entidades.


Resumo do fluxo da API:


### Fluxo da API por caso de uso

---

#### Criar Categoria (POST /categories)

```
[ HTTP Controller: CreateCategory ]
        ↓
[ Use Case: CreateCategoryUseCase ]
        ↓
[ ICategoryRepository.Save() ]
        ↓
[ CategoryRepository (GORM) ]
        ↓
[ MySQL ]
```

---

#### Listar Categorias (GET /categories)

```
[ HTTP Controller: GetCategories ]
        ↓
[ Use Case: GetCategoriesUseCase ]
        ↓
[ ICategoryRepository.FindAll() ]
        ↓
[ CategoryRepository (GORM) ]
        ↓
[ MySQL ]
```

---

#### Buscar Categoria por ID (GET /categories/:id)

```
[ HTTP Controller: GetCategory ]
        ↓
[ Use Case: GetCategoryUseCase ]
        ↓
[ ICategoryRepository.FindById() ]
        ↓
[ CategoryRepository (GORM) ]
        ↓
[ MySQL ]
```

---

#### Atualizar Categoria (PUT /categories/:id)

```
[ HTTP Controller: UpdateCategory ]
        ↓
[ Use Case: UpdateCategoryUseCase ]
        ↓
[ ICategoryRepository.FindById() + FindByName() + Update() ]
        ↓
[ CategoryRepository (GORM) ]
        ↓
[ MySQL ]
```

---

#### Deletar Categoria (DELETE /categories/:id)

```
[ HTTP Controller: DeleteCategory ]
        ↓
[ Use Case: DeleteCategoryUseCase ]
        ↓
[ ICategoryRepository.Delete() ]
        ↓
[ CategoryRepository (GORM) ]
        ↓
[ MySQL ]
```
