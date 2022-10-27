package main

import (
	"real_quote_server/src/application/routes"
	"real_quote_server/src/shared/config"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	if err := createSchema(config.Conn()); err != nil {
		panic(err)
	}

	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	routes.Routes(app)

	app.Run(":8080")
}

func createSchema(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS quote (bid INT NOT NULL)")
	if err != nil {
		return err
	}

	return nil
}
