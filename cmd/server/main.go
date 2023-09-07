package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Waelson/internal/controller"
	"github.com/Waelson/internal/dao"
	"github.com/Waelson/internal/database"
	"github.com/Waelson/internal/repository"
	"github.com/Waelson/internal/service"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ContextTimeoutDatabase = 200 * time.Millisecond
	ContextTimeoutApi      = 10 * time.Millisecond
	EndpointCotacao        = "/cotacao"
	Porta                  = ":8080"
)

func main() {
	fmt.Println("### SERVIDOR DE COTACAO DE DOLAR ###")
	fmt.Println("Obtendo conexao com o banco de dados")

	//Conecta ao SQLite
	db, err := database.ObterConexao()
	if err != nil {
		fmt.Println("nao foi possivel obter a conexao com o banco de dados")
		panic(err)
	}

	fmt.Println("Inicializando a tabela de cotacao")

	//Cria a tabela de cotacao, caso nao exista
	err = database.Inicializar(db)
	if err != nil {
		fmt.Println("ocorreu um erro ao executar o script de criacao da tabela de cotacao")
		panic(err)
	}

	fmt.Println("Criando as dependencias")

	//Inicializa as dependencias
	cotacaoDao := dao.NewCotacaoDao(ContextTimeoutApi)
	cotacaoRepository := repository.NewCotacaoRepository(db, ContextTimeoutDatabase)
	cotacaoService := service.NewCotacaoService(cotacaoDao, cotacaoRepository)
	cotacaoController := controller.NewCotacaoController(cotacaoService)

	fmt.Println("Registrando a rota")

	//Registra a rota
	http.HandleFunc(EndpointCotacao, cotacaoController.Handle)

	fmt.Println("Servidor inicializado na porta", Porta)

	//Inicializa o servidor HTTP
	err = http.ListenAndServe(Porta, nil)
	if err != nil {
		panic(err)
	}

}
