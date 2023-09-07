package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	UrlServidorLocal = "http://localhost:8080/cotacao"
	CaminhoArquivo   = "cotacao.txt"
	TimeoutRequest   = 300 * time.Millisecond
)

func main() {
	fmt.Println("Obtendo a cotacao na URL ", UrlServidorLocal)
	cotacao, err := obterCotacao()
	if err != nil {
		fmt.Println("Ocorreu um erro ao obter a cotacao")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Salvando a cotacao '%s' no arquivo '%s'", cotacao, CaminhoArquivo))

	cotacao = fmt.Sprintf("[%s] %s", obterDataHoraAtual(), cotacao)

	err = salvarCotacao(cotacao)
	if err != nil {
		fmt.Println("Ocorreu um erro ao salvar a cotacao")
		panic(err)
	}

	fmt.Println("Cotacao salva com sucesso")
}

func obterCotacao() (string, error) {
	req, err := http.NewRequest("GET", UrlServidorLocal, nil)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), TimeoutRequest)
	defer cancel()

	// Associe o contexto à solicitação
	req = req.WithContext(ctx)

	// Faz a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func obterDataHoraAtual() string {
	horaAtual := time.Now()
	formatoDesejado := "02/01/2006 15:04:05"
	dataFormatada := horaAtual.Format(formatoDesejado)
	return dataFormatada
}

func criarArquivoSeNaoExiste() (*os.File, error) {
	arquivo, err := os.OpenFile(CaminhoArquivo, os.O_WRONLY|os.O_APPEND, 0644)
	if os.IsNotExist(err) {
		arquivo, err = os.Create(CaminhoArquivo)
		if err != nil {
			return nil, err
		}
		return arquivo, nil
	} else if err != nil {
		return nil, err
	} else {
		return arquivo, nil
	}
}

func salvarCotacao(cotacao string) error {
	arquivo, err := criarArquivoSeNaoExiste()
	if err != nil {
		return err
	}
	nb, err := arquivo.WriteString(fmt.Sprintf("%s\n", cotacao))
	if err != nil {
		return err
	}
	fmt.Printf("Foram escritos %d bytes no arquivo %s\n", nb, CaminhoArquivo)
	return nil
}
