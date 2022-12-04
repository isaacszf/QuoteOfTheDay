# QuoteOfTheDay

[@SomeDaiIyQuotes](https://twitter.com/SomeDaiIyQuotes) is a simple Twitter bot that tweets quotes from famous authors daily. This is mainly done using a
web scrapper on the website [BrainyQuote](https://www.brainyquote.com/quote_of_the_day) and then parsing it so that the quote can be tweeted.

## Running

To be able to run this project, you will need a [developer account](https://developer.twitter.com/en/portal/dashboard) and a project created in it. Then, you
will get the project credentials and place it in order inside the ".env" file.

After that, run:
```bash
go build -o out ./app/.

# --time is optional
./out --time 03:00
```

## Technologies

Packages that were used for this project:

- **[gocron](https://github.com/go-co-op/gocron)**
- **[goquery](https://github.com/PuerkitoBio/goquery)**
- **[godotenv](https://github.com/joho/godotenv)**
- **[gotwitter-v2](https://github.com/g8rswimmer/go-twitter) & [oauth1](https://github.com/dghubble/oauth1)**

To keep the bot always online, [Railway](https://render.com/) was used. 
