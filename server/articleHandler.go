package server

import (
	"google.golang.org/appengine"
	"net/http"
	"encoding/json"
	"github.com/BloodyRainer/articlePrice/search"
	engLog "google.golang.org/appengine/log"
	"github.com/BloodyRainer/articlePrice/dialogflow"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	ctx := appengine.NewContext(req)

	if req.Method == http.MethodPost {
		logPostRequest(ctx, req)
	}

	//TODO: GetArctileByNumber if request contains ArticleNr
	a, err := search.GetRandomArticle(req)
	if err != nil {
		engLog.Errorf(ctx, err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	dr := dialogflow.MakeArticleNameResponse(*a)

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(dr)
	//res.Write([]byte(justTrying5()))

	//engLog.Infof(ctx, "response:", dr.Payload.Google.RichResponse.Items[0].SimpleResponse.TextToSpeech)

}

func justTrying() string {
	return `{"conversationToken":"[]","expectUserResponse":true,"expectedInputs":[{"inputPrompt":{"richInitialPrompt":{"items":[{"simpleResponse":{"textToSpeech":"Hallo, willkommen zu der Preis ist heiß. Sage 'nächster Artikel' um zu starten!"}}]}},"possibleIntents":[{"intent":"assistant.intent.action.TEXT"}]}],"responseMetadata":{"status":{"message":"Success (200)"},"queryMatchInfo":{"queryMatched":true,"intent":"9230d0d0-f738-439d-ab10-b1d75ab9e477"}}}`
}

func justTrying2() string {
	return `{"fulfillmentText":"Thisisatextresponse","fulfillmentMessages":[{"card":{"title":"cardtitle","subtitle":"cardtext","imageUri":"https://assistant.google.com/static/images/molecule/Molecule-Formation-stop.png","buttons":[{"text":"buttontext","postback":"https://assistant.google.com/"}]}}],"source":"example.com","payload":{"google":{"expectUserResponse":true,"richResponse":{"items":[{"simpleResponse":{"textToSpeech":"thisisasimpleresponse"}}]}},"facebook":{"text":"Hello,Facebook!"},"slack":{"text":"ThisisatextresponseforSlack."}},"outputContexts":[{"name":"projects/${PROJECT_ID}/agent/sessions/${SESSION_ID}/contexts/contextname","lifespanCount":5,"parameters":{"param":"paramvalue"}}],"followupEventInput":{"name":"eventname","languageCode":"en-US","parameters":{"param":"paramvalue"}}`
}

func justTrying3() string {
	return `{"conversationToken":"","expectUserResponse":true,"expectedInputs":[{"inputPrompt":{"richInitialPrompt":{"items":[{"simpleResponse":{"textToSpeech":"Howdy!Icantellyoufunfactsaboutalmostanynumber,like42.Whatdoyouhaveinmind?","displayText":"Howdy!Icantellyoufunfactsaboutalmostanynumber.Whatdoyouhaveinmind?"}}],"suggestions":[]}},"possibleIntents":[{"intent":"actions.intent.TEXT"}]}]}`
}

func justTrying4() string {
	return `{ source: 'source-of-the-response',speech: 'message read by voice assistant',displayText: 'message displayed on the user device screen.'}`
}

func justTrying5() string {
	return `{"conversationToken":"[]","expectUserResponse":true,"expectedInputs":[{"inputPrompt":{"richInitialPrompt":{"items":[{"simpleResponse":{"textToSpeech":"sinnvoll"}}]}},"possibleIntents":[{"intent":"assistant.intent.action.TEXT"}]}],"responseMetadata":{"status":{"message":"Success(200)"},"queryMatchInfo":{"queryMatched":true,"intent":"8a757836-9f86-4006-b98d-eeb1e56c051c"}}}`
}
