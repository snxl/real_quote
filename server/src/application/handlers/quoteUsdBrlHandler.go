package handlers

import (
	"net/http"

	quotesclients "real_quote_server/src/core/clients/quotes_clients"
	quoterepository "real_quote_server/src/core/repositories/quote_repository"
	consultrealdolarquoteusecase "real_quote_server/src/core/usecases/consult_real_dolar_quote_usecase"
	"real_quote_server/src/shared/config"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type QuoteUsdBrlHandler struct {
	conn *sqlx.DB
}

func NewQuoteUsdBrlHandler() *QuoteUsdBrlHandler {
	db := config.Conn()

	return &QuoteUsdBrlHandler{db}
}

func (q *QuoteUsdBrlHandler) Handle(c *gin.Context) {
	usecase := consultrealdolarquoteusecase.NewConsultRealDolarQuote(
		quotesclients.NewQuotesclientsEconomiaAwesomeapi(),
		quoterepository.NewQuoteSqlxRepository(q.conn),
	)

	output, err := usecase.Run()
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bid": output.Bid,
	})
}
