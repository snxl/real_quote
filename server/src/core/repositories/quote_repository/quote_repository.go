package quoterepository

import (
	"context"
	"real_quote_server/src/core/models"
)

type QuoteRepository interface {
	CreteQuoteWithContext(ctx context.Context, quote models.Quote) error
}
