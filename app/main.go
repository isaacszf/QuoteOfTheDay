package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"quoteoftheday.isaacszf.net/app/quote"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("🚀 App Started!\n\n")

	// Scheduler
	schedule := flag.Bool("schedule", false, "Schedule Post")

	defaultTime := loadEnvKey("SCHEDULER_TIME")
	schedulerTime := flag.String("time", defaultTime, "Scheduler Time")

	flag.Parse()

	if *schedule {
		scheduler := gocron.NewScheduler(time.UTC)

		scheduler.Every(1).Day().At(*schedulerTime).Do(func() {
			sendTweet()
		})

		scheduler.StartBlocking()
	} else {
		sendTweet()
		os.Exit(1);
	}
}

func sendTweet() {
	fullQuote, err := quote.Load()
	if err != nil {
			log.Fatal(err)
	}

	emoji := getRandomEmoji()
	quote := fmt.Sprintf(`"%s" - %s %s`, fullQuote.Phrase, fullQuote.Author, emoji)

	status := handleTweet(quote)
	log.Print(status)
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
