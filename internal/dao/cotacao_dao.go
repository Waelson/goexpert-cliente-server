package dao

import (
	"context"
	"encoding/json"
	"github.com/Waelson/internal/resource"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

type CotacaoDao interface {
	Obter() (*resource.CotacaoResponse, error)
}

type cotacaoDao struct {
	contextTimeout time.Duration
}

func (e *cotacaoDao) Obter() (*resource.CotacaoResponse, error) {
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	// Associe o contexto à solicitação
	req = req.WithContext(ctx)

	// Faz a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result resource.CotacaoResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewCotacaoDao(contextTimeout time.Duration) CotacaoDao {
	return &cotacaoDao{
		contextTimeout: contextTimeout,
	}
}
