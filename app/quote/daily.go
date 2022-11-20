package quote

var (
	url = "https://www.brainyquote.com/quote_of_the_day/"
)

type Quote struct {
	Phrase string
	Author string
}

func Load() (*Quote, error) {
	document, err := getPageByUrlAndParse(url)
	if err != nil {
		return nil, err
	}

	quote := getFirstElementByClass(document, ".bqQt a div")
	author := getFirstElementByClass(document, ".bq-aut")

	return &Quote{
		Phrase: quote,
		Author: author,
	}, nil
}
