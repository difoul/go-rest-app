package quotes

import (
	"errors"
	"math/rand"
)

type Quote struct {
	Author        string `json:"author"`
	Quote_message string `json:"quote"`
}

func RandomQuote() Quote {
	quotes := AllQuotes()
	return quotes[rand.Intn(len(quotes))]
}

func QuoteById(id int) (Quote, error) {
	quotes := AllQuotes()
	if id >= len(quotes) || id < 0 {
		return Quote{}, errors.New("Not found!")
	}
	return quotes[id], nil
}
