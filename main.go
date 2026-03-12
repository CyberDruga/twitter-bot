package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"
	"syscall"
	"time"

	_ "github.com/CyberDruga/twitter-bot/src/args"
	"github.com/CyberDruga/twitter-bot/src/cache"
	"github.com/CyberDruga/twitter-bot/src/config"
	"github.com/CyberDruga/twitter-bot/src/discord"
	_ "github.com/CyberDruga/twitter-bot/src/logger"
	"github.com/CyberDruga/twitter-bot/src/models"
	"github.com/CyberDruga/twitter-bot/src/trap"
	"github.com/gorilla/websocket"
)

const (
	URL        = "wss://ws.twitterapi.io/twitter/tweet/websocket"
	CACHE_FILE = ".cache"
)

func main() {

	conf, err := config.LoadConfig("./config.toml")

	if err != nil {
		panic("Error: " + err.Error())
	}

	if err := cache.LoadCache(".cache"); err != nil {
		panic("Error: " + err.Error())
	}

	trap.Trap(func() { cache.SaveCache(CACHE_FILE) }, syscall.SIGTERM, syscall.SIGINT)

	headers := http.Header{}

	headers.Add("x-api-key", conf.ApiToken)

	con, _, err := websocket.DefaultDialer.Dial(URL, headers)

	if err != nil {
		panic("Error: " + err.Error())
	}

	for {
		_, bytes, err := con.ReadMessage()

		if err != nil {
			slog.Error("Error: " + err.Error())
			os.Exit(1)
		}

		var message models.WebsocketMessage

		json.Unmarshal(bytes, &message)

		slog.Info(fmt.Sprintf("%s: %v - tweets(%d)", message.RuleId, message.EventType, len(message.Tweets)))

		if message.EventType != "tweet" {
			continue
		}

		if len(message.Tweets) == 0 {
			continue
		}

		slices.Reverse(message.Tweets)

		slog.Debug(fmt.Sprintf("Rules to process: %d", len(conf.Rules)))

		for _, rule := range conf.Rules {
			slog.Debug("Processing rules")
			go HandleTweets(rule, message)
		}

	}
}

func HandleTweets(rule config.Rule, message models.WebsocketMessage) {

	var err error

	slog.Debug("Processing rule" + message.RuleId)
	defer slog.Debug(fmt.Sprintf("End processing rule. Error: %v", err))

	if message.RuleId != rule.RuleId {
		return
	}

	first := true

	cache.Lock()
	defer cache.Unlock()

	for _, tweet := range message.Tweets {

		if slices.Contains(*cache.Tweets, tweet) {
			slog.Debug("Found in cache, ignoring")
			continue
		}

		if first && rule.Message != "" {
			err = discord.SendWebhookMessage(rule.WebhookUrl, rule.Message)

			if err != nil {
				slog.Error("Error: " + err.Error())
			}

			first = false
		}

		msg := strings.Replace(tweet.Url, "x.com", "fixupx.com", 1)

		err = discord.SendWebhookMessage(rule.WebhookUrl, msg)

		cache.AddTweet(tweet)

		time.Sleep(1 * time.Second)

		if err != nil {
			slog.Error("Error: " + err.Error())
		}

	}

}
