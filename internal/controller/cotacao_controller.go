package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Waelson/internal/service"
)

type CotacaoController interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type cotacaoController struct {
	cotacaoService service.CotacaoService
}

func (c *cotacaoController) Handle(w http.ResponseWriter, r *http.Request) {
	cotacao, err := c.cotacaoService.Obter(context.Background())
	if err != nil {
		fmt.Println("ocorreu um erro no processamento da requisicao", err)
		fmt.Fprintf(w, "Erro ao processar a requisicao")
		return
	}
	fmt.Println("Controller sucesso...", cotacao)
	fmt.Fprintf(w, "%s", cotacao)
}

func NewCotacaoController(cotacaoService service.CotacaoService) CotacaoController {
	controller := cotacaoController{
		cotacaoService: cotacaoService,
	}
	return &controller
}
