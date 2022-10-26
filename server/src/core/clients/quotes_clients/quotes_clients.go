package quotesclients

import "context"

type Quotesclients interface {
	ConsultBrlUsdQuoteWithContext(ctx context.Context) (*ConsultBrlUsdQuoteOutput, error)
}

type ConsultBrlUsdQuoteOutput struct {
	Bid string
}
