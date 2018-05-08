package model

import "fmt"

type Article struct {
	ArticleNr string `json:"articleNumber"`
	Name      string `json:"articleName"`
	Price     string `json:"articlePrice"`
	ImgUrl    string `json:"imgUrl"`
}

func (rcv Article) String() string {
	return fmt.Sprintf("[ArticleNr: %v, Name: %v, Price: %v, ImgUrl: %v]", rcv.ArticleNr, rcv.Name, rcv.Price, rcv.ImgUrl)
}
