package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/charmbracelet/log"
	"io"
	"net/http"
)

type DiscordMessage struct {
	Content string `json:"content,omitempty"`
}

func SendWebhookMessage(url string, message string) (err error) {

	discordMessage := DiscordMessage{
		Content: message,
	}

	body, err := json.Marshal(discordMessage)

	if err != nil {
		return
	}

	log.Info("Sending message", "Message", message)

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
		return errors.New(response.Message)
	}

	return
}
