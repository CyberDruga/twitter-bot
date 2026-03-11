package models

type WebsocketMessage struct {
	EventType string  `json:"event_type"`
	RuleId    string  `json:"rule_id"`
	Tweets    []Tweet `json:"tweets"`
}

type Tweet struct {
	Url string `json:"url"`
}
