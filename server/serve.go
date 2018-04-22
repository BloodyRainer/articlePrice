package server

import (
	"net/http"
	"github.com/BloodyRainer/articlePrice/search"
	"encoding/json"
	"log"
	"io/ioutil"
	"fmt"
)

type articleHandler struct{}

func (rcv *articleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	a, err := search.GetRandomArticle()
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	req.Header.Set("Content-Type", "application/json")

	enc := json.NewEncoder(res)
	enc.Encode(a)

	log.Println("responded Article:", *a)

	if req.Method == http.MethodPost {
		printPostRequest(*req)
	}

}

func StartServer() {
	log.Println("Starting Server...")
	http.Handle("/getArticle", &articleHandler{})
	http.ListenAndServe(":8444", nil)
}

func printPostRequest(req http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(string(body))

}
