package dialogflow

import "encoding/json"

type Message struct {
	Card            *Card            `json:"card,omitempty"`
	Platform        string           `json:"platform,omitempty"`
	SimpleResponses *SimpleResponses `json:"simpleResponses,omitempty"`
	Text            *Text            `json:"text,omitempty"`
}

type Text struct {
	Text []string `json:"text,omitempty"`
}

type Context struct {
	Name          string          `json:"name,omitempty"`
	LifespanCount int             `json:"lifespanCount,omitempty"`
	Parameters    json.RawMessage `json:"parameters,omitempty"`
}
