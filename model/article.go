package model

import "fmt"

type Article struct {
	ArticleNr string `json: articleNr`
	Name      string `json: name`
	Price     string `json: price`
}

func (rcv Article) String() string {
	return fmt.Sprintf("[ArticleNr: %v, Name: %v, Price: %v]", rcv.ArticleNr, rcv.Name, rcv.Price)
}