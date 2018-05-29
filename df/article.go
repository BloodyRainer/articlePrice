package df

import (
	"fmt"
)

type Article struct {
	ArticleNr string `json:"articleNumber"`
	Name      string `json:"articleName"`
	Price     string `json:"articlePrice"`
	ImgUrl    string `json:"imgUrl"`
	Link      string `json:"link"`
}

func (rcv Article) String() string {
	return fmt.Sprintf("[ArticleNr: %v, Name: %v, Price: %v, ImgUrl: %v, Link: %v]", rcv.ArticleNr, rcv.Name, rcv.Price, rcv.ImgUrl, rcv.Link)
}

func (rcv Article) ToParameters() []byte {
	params := MakeParameters("articleNumber", rcv.ArticleNr)
	params = AppendString(params, "articleName", rcv.Name)
	params = AppendNonString(params, "actualPrice", rcv.Price)
	params = AppendString(params, "imgUrl", rcv.ImgUrl)
	params = AppendString(params, "link", rcv.Link)

	return params
}
