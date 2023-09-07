package repository

import (
	"context"
	"database/sql"
	"time"
)

const (
	sqlInsert = "INSERT into cotacao (valor) VALUES (?)"
)

type CotacaoRepository interface {
	Salvar(cotacao float64) error
}

type cotacaoRepository struct {
	db             *sql.DB
	contextTimeout time.Duration
}

func (c *cotacaoRepository) Salvar(cotacao float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.contextTimeout)
	defer cancel()
	_, err := c.db.ExecContext(ctx, sqlInsert, cotacao)
	return err
}

func NewCotacaoRepository(db *sql.DB, contextTimeout time.Duration) CotacaoRepository {
	return &cotacaoRepository{
		db:             db,
		contextTimeout: contextTimeout,
	}
}
