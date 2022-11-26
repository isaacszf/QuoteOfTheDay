package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

	"quoteoftheday.isaacszf.net/app/quote"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("🚀 App Started!")

	// For render
	mux := http.NewServeMux()
	mux.HandleFunc("/health_check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("App is running!"))
	})
	defer http.ListenAndServe(":10000", mux)

	// Scheduler
	schedulerTime := flag.String("time", loadEnvKey("SCHEDULER_TIME"), "Scheduler Time")
	flag.Parse()

	scheduler := gocron.NewScheduler(time.UTC)

	log.Println("📅 Scheduler Created..")

	scheduler.Every(1).Day().At(*schedulerTime).Do(func() {
		fullQuote, err := quote.Load()
		if err != nil {
			log.Fatal(err)
		}

		emoji := getRandomEmoji()
		quote := fmt.Sprintf(`"%s" - %s %s`, fullQuote.Phrase, fullQuote.Author, emoji)

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
