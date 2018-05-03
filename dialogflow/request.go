package dialogflow

import (
	"encoding/json"
	"errors"
)

func MakeDfRequest(body []byte) (*DfRequest, error) {

	var dfRequest DfRequest

	err := json.Unmarshal(body, &dfRequest)
	if err != nil {
		return nil, errors.New("failed to parse http request into dialogflow request: " + err.Error())
	}

	return &dfRequest, nil

}

type DfRequest struct {
	ResponseId  string       `json:"responseId,omitempty"`
	Session     string       `json:"session,omitempty"`
	QueryResult *QueryResult `json:"queryResult,omitempty"`
}

type QueryResult struct {
	QueryText                   string          `json:"queryText,omitempty"`
	Parameters                  json.RawMessage `json:"parameters,omitempty"`
	AllRequiredParamsPresent    bool            `json:"allRequiredParamsPresent,omitempty"`
	FulfillmentText             string          `json:"fulfillmentText,omitempty"`
	FulfillmentMessages         []Message       `json:"fulfillmentMessages,omitempty"`
	OutputContexts              []Context       `json:"outputContexts,omitempty"`
	Intent                      *Intent         `json:"intent,omitempty"`
	IntentDetectionConfidence   int             `json:"intentDetectionConfidence,omitempty"`
	DiagnosticInfo              *DiagnosticInfo `json:"diagnosticInfo,omitempty"`
	LanguageCode                string          `json:"languageCode,omitempty"`
	OriginalDetectIntentRequest DfRequest       `json:"originalDetectIntentRequest,omitempty"`
}

type Intent struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type DiagnosticInfo struct {
}
