package quoterepository

import (
	"context"
	"real_quote_server/src/core/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type QuoteSqlxRepository struct {
	conn *sqlx.DB
}

func NewQuoteSqlxRepository(conn *sqlx.DB) *QuoteSqlxRepository {
	return &QuoteSqlxRepository{conn: conn}
}

func (q *QuoteSqlxRepository) CreteQuoteWithContext(ctx context.Context, quote models.Quote) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	stmt, err := q.conn.PrepareContext(ctx, "INSERT INTO quote(bid) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(quote.Bid)
	if err != nil {
		return err
	}

	return nil
}
