package df

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

type Payload struct {
	Google            *Google       `json:"google,omitempty"`
	IsInSandbox       bool          `json:"isInSandbox,omitempty"`
	Surface           *Surface      `json:"surface,omitempty"`
	Inputs            []Input       `json:"inputs,omitempty"`
	User              *User         `json:"user,omitempty"`
	Conversation      *Conversation `json:"conversation,omitempty"`
	AvailableSurfaces []Surface     `json:"availableSurfaces,omitempty"`
}

type Conversation struct {
	ConversationId    string `json:"conversationId,omitempty"`
	Type              string `json:"type,omitempty"`
	ConversationToken string `json:"conversationToken,omitempty"`
}

type User struct {
	LastSeen string `json:"lastSeen,omitempty"`
	Locale   string `json:"locale,omitempty"`
	UserId   string `json:"userId,omitempty"`
}

type Surface struct {
	Capabilities []Capability `json:"capabilities,omitempty"`
}

type Capability struct {
	Name string `json:"name,omitempty"`
}

type Input struct {
	RawInputs []RawInput `json:"rawinputs,omitempty"`
	Arguments []Argument `json:"arguments,omitempty"`
	Intent    string     `json:"intent,omitempty"`
}

type RawInput struct {
	Query     string `json:"query,omitempty"`
	InputType string `json:"inputType,omitempty"`
}

type Argument struct {
	RawText   string `json:"rawText,omitempty"`
	TextValue string `json:"textValue,omitempty"`
	Name      string `json:"name,omitempty"`
}
