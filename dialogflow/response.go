package dialogflow

import (
	"github.com/BloodyRainer/articlePrice/model"
	"encoding/json"
)

type DfResponse struct {
	FulfillmentText     string      `json:"fulfillmentText,omitempty"`
	FulfillmentMessages []Message   `json:"fulfillmentMessages,omitempty"`
	Source              string      `json:"source,omitempty"`
	Payload             *Payload    `json:"payload,omitempty"` //TODO: is this actually called data? https://developers.google.com/actions/dialogflow/webhook
	OutputContexts      []Context   `json:"outputContexts,omitempty"`
	FollowupEventInput  *EventInput `json:"followupEventInput,omitempty"`
}

func MakeArticleNameResponse(a model.Article) *DfResponse {
	dr := &DfResponse{
		Source: "Der Preis ist heiss",
		Payload: &Payload{
			Google: &Google{
				ExpectUserResponse: true,
				RichResponse: &RichResponse{
					Items: []Item{
						{
							SimpleResponse: &SimpleResponse{
								TextToSpeech: "Wie ist der Preis von: " + a.Name + "?",
							},
						},
					},
				},
			},
		},
	}
	return dr
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
	Items []Item `json:"items,omitempty"`
}

type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
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
