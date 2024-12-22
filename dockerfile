# Use uma imagem base leve com suporte ao Go
FROM golang:1.23.4

# Defina o diretório de trabalho
WORKDIR /app

# Copie go.mod e go.sum primeiro
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copie o restante dos arquivos do projeto
COPY . .

# Especifique o caminho do pacote principal para o build
RUN go build -o stress-tester ./cmd/stress-tester

# Comando de execução
ENTRYPOINT ["./stress-tester"]
