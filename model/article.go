package model

import "fmt"

type Article struct {
	ArticleNr string `json: articleNumber`
	Name      string `json: articleName`
	Price     string `json: articlePrice`
}

func (rcv Article) String() string {
	return fmt.Sprintf("[ArticleNr: %v, Name: %v, PriceGuess: %v]", rcv.ArticleNr, rcv.Name, rcv.Price)
}