package service

import (
	"context"
	"fmt"
	"github.com/Waelson/internal/dao"
	"github.com/Waelson/internal/repository"
	"strconv"
)

type CotacaoService interface {
	Obter(ctx context.Context) (string, error)
}

type cotacaoService struct {
	cotacaoRepository repository.CotacaoRepository
	cotacaoDao        dao.CotacaoDao
}

func (c *cotacaoService) Obter(ctx context.Context) (string, error) {

	fmt.Println("Obtendo a cotacao da API")

	//Obtem a cotacao do dolar via API
	resultadoApi, err := c.cotacaoDao.Obter()
	if err != nil {
		return "", err
	}

	valorDolar, err := strconv.ParseFloat(resultadoApi.USDBRL.Bid, 64)
	if err != nil {
		return "", err
	}

	fmt.Println("Persistindo a cotacao no SQLite")

	//Persiste a cotacao no SQLite
	err = c.cotacaoRepository.Salvar(valorDolar)
	if err != nil {
		return "", err
	}

	return resultadoApi.USDBRL.Bid, nil
}

func NewCotacaoService(cotacaoDao dao.CotacaoDao,
	cotacaoRepository repository.CotacaoRepository) CotacaoService {
	return &cotacaoService{
		cotacaoRepository: cotacaoRepository,
		cotacaoDao:        cotacaoDao,
	}
}
