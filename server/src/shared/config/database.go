package config

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var lock = &sync.Mutex{}
var sqlxConn *sqlx.DB

func Conn() *sqlx.DB {
	if sqlxConn == nil {
		lock.Lock()
		defer lock.Unlock()
		if sqlxConn == nil {
			db, err := sqlx.Open("sqlite3", ":memory:")
			if err != nil {
				panic(err)
			}

			sqlxConn = db
		}
	}

	return sqlxConn
}
