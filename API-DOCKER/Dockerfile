# Etapa 1: build da aplicação
FROM golang:1.18 AS builder

# Define o diretório de trabalho
WORKDIR /go/src/app

# Copia os arquivos para o container
COPY . .

# Baixa as dependências (opcional, mas acelera builds)
RUN go mod download

# Compila o binário estaticamente
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Etapa 2: imagem final minimalista
FROM scratch

# Copia o binário gerado na etapa anterior
COPY --from=builder /go/src/app/main /main

# Expõe a porta da aplicação (se necessário)
EXPOSE 8000

# Comando de entrada
ENTRYPOINT ["/main"]

