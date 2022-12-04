package main

import (
	"fmt"
	"os"
	"time"

	"github.com/slack-go/slack"
	"github.com/your/project/parse_stock_ticker"
	"github.com/autotune/stocksgpt/slackbot"
)

func main() {
	ticker := os.Getenv("STOCK_TICKER")
	token := os.Getenv("SLACK_TOKEN")
	client := slack.New(token)

	for {
		price, err := parse_stock_ticker.ParseStockTicker(ticker)
		if err != nil {
			fmt.Printf("Error getting stock price: %s\n", err)
			break
		}

		if price >= 1 {
			slackbot.PostToSlack(client, "Finance", fmt.Sprintf("%s stock has moved up $1", ticker))
		} else if price <= -1 {
			slackbot.PostToSlack(client, "Finance", fmt.Sprintf("%s stock has moved down $1", ticker))
		}

		time.Sleep(3 * time.Second)
	}
}
