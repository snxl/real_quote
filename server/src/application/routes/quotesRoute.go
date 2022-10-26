package routes

import (
	"real_quote_server/src/application/handlers"

	"github.com/gin-gonic/gin"
)

func quotesRoutes(eng *gin.Engine) {
	eng.GET("/cotacao", handlers.NewQuoteUsdBrlHandler().Handle)
}
