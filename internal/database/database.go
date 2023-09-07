package database

import (
	"database/sql"
)

const (
	SqlCreateTable = "CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY AUTOINCREMENT, valor REAL)"
	CamindoDB      = "./db/cotacao.db"
	DriverDb       = "sqlite3"
)

func ObterConexao() (*sql.DB, error) {
	db, err := sql.Open(DriverDb, CamindoDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func Inicializar(db *sql.DB) error {
	_, err := db.Exec(SqlCreateTable)
	if err != nil {
		return err
	}
	return nil
}
