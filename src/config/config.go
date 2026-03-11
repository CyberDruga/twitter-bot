package config

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ApiToken string `toml:"api_token"`
	Rules    []Rule `toml:"rule"`
}

type Rule struct {
	WebhookUrl string `toml:"webhook_url"`
	RuleId     string `toml:"rule_id"`
	Message    string `toml:"message"`
}

func LoadConfig(path string) (config Config, err error) {

	slog.Debug("Loading config")
	defer slog.Debug(fmt.Sprintf("Done loading config. Error: %v", err))

	if _, err = toml.DecodeFile(path, &config); err != nil {
		return
	}

	if config.ApiToken == "" {
		err = errors.New("No api token informed")
		return
	}

	if len(config.Rules) == 0 {
		err = errors.New("No rules configured")
		return
	}

	for _, rule := range config.Rules {
		if rule.RuleId == "" {
			err = errors.New("No rule_id informed")
			return
		}

		if rule.WebhookUrl == "" {
			err = errors.New("No webhook url informed")
			return
		}
	}

	return
}
