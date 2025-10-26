package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {

	db, err := sql.Open("pgx", config.UrlDatabase)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
