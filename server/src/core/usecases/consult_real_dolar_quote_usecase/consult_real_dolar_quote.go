package consultrealdolarquoteusecase

import (
	"context"
	"net/http"
	quotesclients "real_quote_server/src/core/clients/quotes_clients"
	"real_quote_server/src/core/models"
	quoterepository "real_quote_server/src/core/repositories/quote_repository"
	"real_quote_server/src/shared/helper"
)

type ConsultRealDolarQuote struct {
	quotesClient quotesclients.Quotesclients
	quoteRepo    quoterepository.QuoteRepository
}

func NewConsultRealDolarQuote(
	quotesClient quotesclients.Quotesclients,
	quoteRepo quoterepository.QuoteRepository,
) *ConsultRealDolarQuote {
	return &ConsultRealDolarQuote{
		quotesClient: quotesClient,
		quoteRepo:    quoteRepo,
	}
}

func (c *ConsultRealDolarQuote) Run() (*models.Quote, *helper.HttpError) {

	ctx := context.Background()

	getQuote, err := c.quotesClient.ConsultBrlUsdQuoteWithContext(ctx)
	if err != nil {
		return nil, &helper.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err,
		}
	}

	quote := models.Quote{
		Bid: getQuote.Bid,
	}

	err = c.quoteRepo.CreteQuoteWithContext(ctx, quote)
	if err != nil {
		return nil, &helper.HttpError{
			Status:  http.StatusInternalServerError,
			Message: err,
		}
	}

	return &quote, nil
}
