package config

import (
	"errors"
	"github.com/charmbracelet/log"

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

	log.Debug("Loading config")
	defer log.Debug("Done loading config", "Error", err)

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
