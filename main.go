package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/websocket"
)

const (
	URL = "wss://ws.twitterapi.io/twitter/tweet/websocket"
)

type Config struct {
	WebhookUrl string `toml:"webhook_url"`
	RuleId     string `toml:"rule_id"`
	ApiToken   string `toml:"api_token"`
	Message    string `toml:"message"`
}

type WebhookMessage struct {
	EventType string  `json:"event_type"`
	RuleId    string  `json:"rule_id"`
	Tweets    []Tweet `json:"tweets"`
}

type Tweet struct {
	Url string `json:"url"`
}

func main() {

	var config Config

	toml.DecodeFile("./config.toml", &config)

	if config.ApiToken == "" {
		panic("No api token informed")
	}

	if config.RuleId == "" {
		panic("No rule_id informed")
	}

	if config.WebhookUrl == "" {
		panic("No webhook url informed")
	}

	headers := http.Header{}

	headers.Add("x-api-key", config.ApiToken)

	con, _, err := websocket.DefaultDialer.Dial(URL, headers)

	if err != nil {
		panic("Error: " + err.Error())
	}

	for {
		_, bytes, err := con.ReadMessage()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: "+err.Error())
			return
		}

		fmt.Println(string(bytes))

		var message WebhookMessage

		json.Unmarshal(bytes, &message)

		if message.EventType != "tweet" {
			continue
		}

		fmt.Printf("%v\n", message)

		slices.Reverse(message.Tweets)

		for _, tweet := range message.Tweets {
			var message string

			if config.Message != "" {
				message = config.Message + " "
			}

			message = message + strings.Replace(tweet.Url, "x.com", "fixupx.com", 1)

			err := SendWebhookMessage(config.WebhookUrl, message)

			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: "+err.Error())
			}

		}

	}
}

func SendWebhookMessage(url string, message string) (err error) {

	type DiscordMessage struct {
		Content string `json:"content,omitempty"`
	}

	discordMessage := DiscordMessage{
		Content: message,
	}

	body, err := json.Marshal(discordMessage)

	if err != nil {
		return
	}

	fmt.Println("sending message")

	res, err := http.Post(url+"?wait=true", "application/json", bytes.NewBuffer(body))

	if err != nil {
		return
	}

	defer res.Body.Close()

	stuff, err := io.ReadAll(res.Body)

	type Response struct {
		Message string `json:"message"`
	}

	var response Response

	err = json.Unmarshal(stuff, &response)

	if err != nil {
		return
	}

	if response.Message != "" {
		return fmt.Errorf(response.Message)
	}

	return
}
