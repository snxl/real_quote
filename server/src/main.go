package main

import (
	"real_quote_server/src/application/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	routes.Routes(app)

	app.Run(":8080")
}
