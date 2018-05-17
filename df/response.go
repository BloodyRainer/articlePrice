package df

import (
	"encoding/json"
)

type Response struct {
	FulfillmentText     string      `json:"fulfillmentText,omitempty"`
	FulfillmentMessages []Message   `json:"fulfillmentMessages,omitempty"`
	Source              string      `json:"source,omitempty"`
	Payload             *Payload    `json:"payload,omitempty"` //TODO: is this actually called data? https://developers.google.com/actions/dialogflow/webhook
	OutputContexts      []Context   `json:"outputContexts,omitempty"`
	FollowupEventInput  *EventInput `json:"followupEventInput,omitempty"`
}

type SimpleResponses struct {
	SimpleResponses []SimpleResponse `json:"simpleResponses,omitempty"`
}

type EventInput struct {
	Name         string          `json:"name,omitempty"`
	LanguageCode string          `json:"languageCode,omitempty"`
	Parameters   json.RawMessage `json:"parameters,omitempty"`
}

type Google struct {
	ExpectUserResponse bool          `json:"expectUserResponse,omitempty"`
	RichResponse       *RichResponse `json:"richResponse,omitempty"`
}

type RichResponse struct {
	Items             []Item             `json:"items,omitempty"`
	Suggestions       []Suggestion       `json:"suggestions,omitempty"`
	LinkOutSuggestion *LinkOutSuggestion `json:"linkOutSuggestion,omitempty"`
}

type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
	BasicCard      *BasicCard      `json:"basicCard,omitempty"`
}

type Suggestion struct {
	Title string `json:"title,omitempty"`
}

type LinkOutSuggestion struct {
	DestinationName string `json:"destinationName,omitempty"`
	Url             string `json:"url,omitempty"`
}

type BasicCard struct {
	Title         string `json:"title,omitempty"`
	Subtitle      string `json:"subtitle,omitempty"`
	FormattedText string `json:"formattedText,omitempty"`
	Image         *Image `json:"image,omitempty"`
}

type Image struct {
	Url               string `json:"url,omitempty"`
	AccessibilityText string `json:"accessibilityText,omitempty"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

type Card struct {
	Title    string   `json:"title,omitempty"`
	SubTitle string   `json:"subTitle,omitempty"`
	ImageUri string   `json:"imageUri,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

type Button struct {
	Text     string `json:"text,omitempty"`
	Postback string `json:"postback,omitempty"`
}
