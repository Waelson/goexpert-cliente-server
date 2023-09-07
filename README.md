# Go Expert - Cliente / Servidor
Projeto cliente / servidor desenvolvido para empregar o uso de: 
- Tratamento de contextos Golang (`context.Context`)
- Requisições HTTP
- Manipulação de banco de dados SQLite
- Manipulação de arquivos

## Baixando as dependências
```
$ go mod tidy
```

## Servidor
Iniciando o servidor

```
$ cd cmd/server
$ go run main.go
```
O arquivo do banco de dados se encontra no arquivo abaixo:
```
db/cotacao.db
```

## Cliente
Iniciando o servidor
```
$ cd cmd/client
$ go run main.go
```

O arquivo contendo as cotações serão armazenados no arquivo abaixo:
```
cmd/client/cotacao.txt
```