package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string
}

func main() {

	client := http.Client{
		Timeout: time.Millisecond * 300,
	}

	req, err := client.Get("http://localhost:8080/cotacao")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(res, &cotacao)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./tmp/cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	price := fmt.Sprintf("Dólar: %s", cotacao.Bid)
	file.Write([]byte(price))
}
