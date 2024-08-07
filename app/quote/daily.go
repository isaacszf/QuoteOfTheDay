package quote

import (
	"math/rand"
)

var (
	url = "https://www.brainyquote.com/quote_of_the_day"
)

type Quote struct {
	Phrase string
	Author string
}

func Load() (*Quote, error) {
	request, err := makeRequest(url)
	if err != nil {
		return nil, err
	}

	document, err := getPageByUrlAndParse(request)
	if err != nil {

		quote, err := loadRandomQuote();
		if err != nil {
			return nil, err
		}

		return quote, nil;
	}

	quote := getFirstElementByClass(document, ".bqQt a div")
	author := getFirstElementByClass(document, ".bq-aut")

	return &Quote{
		Phrase: quote,
		Author: author,
	}, nil
}

func loadRandomQuote() (*Quote, error) {
	quotes := []Quote{
		{"The only way to do great work is to love what you do.", "Steve Jobs"},
		{"In three words I can sum up everything I've learned about life: it goes on.", "Robert Frost"},
		{"To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.", "Ralph Waldo Emerson"},
		{"If you tell the truth, you don't have to remember anything.", "Mark Twain"},
		{"A friend is someone who knows all about you and still loves you.", "Elbert Hubbard"},
		{"To live is the rarest thing in the world. Most people exist, that is all.", "Oscar Wilde"},
		{"Always forgive your enemies; nothing annoys them so much.", "Oscar Wilde"},
		{"Without music, life would be a mistake.", "Friedrich Nietzsche"},
		{"We accept the love we think we deserve.", "Stephen Chbosky"},
		{"It is never too late to be what you might have been.", "George Eliot"},
		{"Success is not final, failure is not fatal: It is the courage to continue that counts.", "Winston Churchill"},
		{"What lies behind us and what lies before us are tiny matters compared to what lies within us.", "Ralph Waldo Emerson"},
		{"Believe you can and you're halfway there.", "Theodore Roosevelt"},
		{"Change the way you look at things and the things you look at change.", "Wayne Dyer"},
		{"The best time to plant a tree was 20 years ago. The second best time is now.", "Chinese Proverb"},
		{"Your time is limited, so don’t waste it living someone else’s life.", "Steve Jobs"},
		{"You miss 100% of the shots you don’t take.", "Wayne Gretzky"},
		{"Whether you think you can or you think you can’t, you’re right.", "Henry Ford"},
		{"The only limit to our realization of tomorrow is our doubts of today.", "Franklin D. Roosevelt"},
		{"Do what you can, with what you have, where you are.", "Theodore Roosevelt"},
		{"The best way to predict the future is to invent it.", "Alan Kay"},
		{"Life is what happens when you’re busy making other plans.", "John Lennon"},
		{"The purpose of our lives is to be happy.", "Dalai Lama"},
		{"Get busy living or get busy dying.", "Stephen King"},
	}

	randomIndex := rand.Intn(len(quotes))
	return &quotes[randomIndex], nil
}
