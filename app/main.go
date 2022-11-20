package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"quoteoftheday.isaacszf.net/app/quote"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	fullQuote, err := quote.Load()
	if err != nil {
		log.Fatal(err)
	}

	emoji := getRandomEmoji()
	quote := fmt.Sprintf(`"%s" - %s %s`, fullQuote.Phrase, fullQuote.Author, emoji)

	// Scheduler
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(1).Day().At("14:30").Do(func() {
		status := handleTweet(quote)
		log.Println(status)
	})

	scheduler.StartBlocking()
}

func handleTweet(text string) string {
	bot := &TwitterBotConfig{
		ConsumerToken:     loadEnvKey("CONSUMER_TOKEN"),
		ConsumerSecret:    loadEnvKey("CONSUMER_SECRET"),
		AccessToken:       loadEnvKey("ACCESS_TOKEN"),
		AccessTokenSecret: loadEnvKey("ACCESS_TOKEN_SECRET"),
	}

	status, err := bot.Tweet(text)
	if err != nil {
		log.Fatal(err)
	}

	return status
}

func loadEnvKey(key string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}