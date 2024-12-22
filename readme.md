# Stress-Tester CLI

## Descrição

Stress-Tester é uma ferramenta de linha de comando (CLI) desenvolvida em Go para realizar testes de carga em serviços web. Ela permite simular requisições HTTP concorrentes e gerar relatórios detalhados sobre o desempenho do serviço.

## Funcionalidades

 - Realiza requisições HTTP concorrentes para a URL especificada.

 - Gera um relatório com:

   - Tempo total de execução.

   - Número total de requisições realizadas.

   - Número de requisições bem-sucedidas (status HTTP 200).

   - Distribuição de outros códigos de status HTTP.

 - Suporte para salvar relatórios em formato JSON.

## Requisitos
 - Docker
 - Go 1.23 ou superior (para desenvolvimento local)

## Instalação e Execução

## 1. Usando Docker ##

**Construir a imagem**
`docker build -t stress-tester .`

**Exemplo de uso com docker-compose**

Se você desejar executar várias configurações simultaneamente, use o arquivo `docker-compose.yml`:
`docker-compose up`

## 2. Execução Local (sem Docker) ##

**Instalar dependências**

Certifique-se de ter o Go instalado e execute:
`go mod download`

**Compilar o binário**
`go build -o stress-tester ./cmd/stress-tester`

**Executar**
`./stress-tester --url=https://www.google.com --requests=100`

## Parâmetros da CLI ##

 - --url (obrigatório): URL do serviço a ser testado.

 - --requests (opcional, padrão: 1): Número total de requisições a serem realizadas.

 - --concurrency (opcional, padrão: 1): Número de requisições concorrentes.

 - --save (opcional): Salva o relatório em um arquivo JSON chamado stress_test_report.json.

 - --detail (opcional): Inclui detalhes de cada requisição no relatório.

## Estrutura do Projeto ##

```plaintext
├── cmd
│   └── stress-tester
│       └── main.go      # Ponto de entrada da aplicação
├── internal
│   ├── config
│   │   └── config.go    # Gerenciamento de configurações da CLI
│   └── tester
│       ├── tester.go    # Execução dos testes de carga
│       └── report.go    # Geração e exibição de relatórios
├── tests                # Testes unitários
├── docker-compose.yml   # Configuração para execução com Docker Compose
├── Dockerfile           # Configuração do contêiner Docker
└── go.mod               # Dependências do projeto
```

## Relatório Gerado ##

O relatório final inclui:

Tempo total: Tempo gasto para executar todas as requisições.

Requests totais: Total de requisições enviadas.

Status 200: Total de requisições com status HTTP 200.

Distribuição de status HTTP: Frequência de outros códigos de status HTTP (ex.: 404, 500).

Detalhes de cada requisição (opcional): Inclui tempo de resposta e código de status para cada requisição.

```plaintext
{
  "total_time": "2.345s",
  "total_requests": 100,
  "successful_200": 98,
  "status_counts": {
    "200": 98,
    "500": 2
  },
  "requests_details": [
    {"status_code": 200, "duration": "45ms"},
    {"status_code": 500, "duration": "60ms"}
  ]
}
```

## Testes ##

Para rodar os testes unitários:

`go test ./...`

## Contribuição ##

Contribuições são bem-vindas! Por favor, envie um Pull Request com suas alterações ou crie uma issue para sugestões e melhorias.