package main

import (
	"context"
	"net/http"

	"github.com/dghubble/oauth1"
	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type TwitterBotConfig struct {
	ConsumerToken     string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func (bot *TwitterBotConfig) Tweet(text string) (string, error) {
	// HTTP client for tweeting
	config := oauth1.NewConfig(bot.ConsumerToken, bot.ConsumerSecret)

	httpClient := config.Client(oauth1.NoContext, &oauth1.Token{
		Token:       bot.AccessToken,
		TokenSecret: bot.AccessTokenSecret,
	})

	client := &twitter.Client{
		Authorizer: &TwitterBotConfig{},
		Client:     httpClient,
		Host:       "https://api.twitter.com",
	}

	// Creating Tweet
	req := twitter.CreateTweetRequest{Text: text}

	if _, err := client.CreateTweet(context.Background(), req); err != nil {
		return "Error Happened", err
	} else {
		return "🔥 Tweet was sent successfully!", nil
	}
}

func (bot *TwitterBotConfig) Add(req *http.Request) {}
