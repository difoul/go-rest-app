package quotes

import "math/rand"

type Quote struct {
	Author        string `json:"author"`
	Quote_message string `json:"quote"`
}

func RandomQuote() Quote {
	quotes := LoadData()
	return quotes[rand.Intn(len(quotes))]
}
