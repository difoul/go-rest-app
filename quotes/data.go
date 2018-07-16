package quotes

var all_quotes []Quote

func AllQuotes() []Quote {
	all_quotes := []Quote{}

	all_quotes = append(all_quotes, Quote{
		Author:        "Albert Einstein",
		Quote_message: "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	})

	all_quotes = append(all_quotes, Quote{
		Author:        "Friedrich Nietzsche",
		Quote_message: "Without music, life would be a mistake.",
	})

	all_quotes = append(all_quotes, Quote{
		Author:        "Mahatma Gandhi",
		Quote_message: "Be the change that you wish to see in the world.",
	})

	all_quotes = append(all_quotes, Quote{
		Author:        "Oscar Wilde",
		Quote_message: "Be yourself; everyone else is already taken.",
	})

	all_quotes = append(all_quotes, Quote{
		Author:        "Niccolo Machiavelli",
		Quote_message: "It is better to be feared than loved, if you cannot be both.",
	})

	return all_quotes
}
