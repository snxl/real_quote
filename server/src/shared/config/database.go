package config

import (
	"fmt"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Conn() *sqlx.DB {
	pwd, _ := filepath.Abs("")
	sqlite3Path := fmt.Sprintf("%s/src/shared/config/sqlite3.db", pwd)

	db, err := sqlx.Open("sqlite3", sqlite3Path)
	if err != nil {
		panic(err)
	}

	return db
}
