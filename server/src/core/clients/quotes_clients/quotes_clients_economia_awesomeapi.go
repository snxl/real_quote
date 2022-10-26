package quotesclients

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type QuotesclientsEconomiaAwesomeapi struct{}

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type ConsultBrlUsdQuoteApiOutput struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

func NewQuotesclientsEconomiaAwesomeapi() *QuotesclientsEconomiaAwesomeapi {
	return &QuotesclientsEconomiaAwesomeapi{}
}

func (q *QuotesclientsEconomiaAwesomeapi) ConsultBrlUsdQuoteWithContext(ctx context.Context) (*ConsultBrlUsdQuoteOutput, error) {

	ctx, calcel := context.WithTimeout(ctx, time.Millisecond*200)
	defer calcel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quote ConsultBrlUsdQuoteApiOutput
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, err
	}

	return &ConsultBrlUsdQuoteOutput{quote.Usdbrl.Bid}, nil
}
