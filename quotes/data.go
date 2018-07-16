package quotes

func LoadData() []Quote {
	quotes_list := make([]Quote, 0)

	quotes_list = append(quotes_list, Quote{
		Author:        "Albert Einstein",
		Quote_message: "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	})

	quotes_list = append(quotes_list, Quote{
		Author:        "Friedrich Nietzsche",
		Quote_message: "Without music, life would be a mistake.",
	})

	quotes_list = append(quotes_list, Quote{
		Author:        "Mahatma Gandhi",
		Quote_message: "Be the change that you wish to see in the world.",
	})

	quotes_list = append(quotes_list, Quote{
		Author:        "Oscar Wilde",
		Quote_message: "Be yourself; everyone else is already taken.",
	})

	quotes_list = append(quotes_list, Quote{
		Author:        "Niccolo Machiavelli",
		Quote_message: "It is better to be feared than loved, if you cannot be both.",
	})

	return quotes_list
}
