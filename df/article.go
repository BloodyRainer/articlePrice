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
	params = AppendParameter(params, "articleName", rcv.Name)
	params = AppendParameter(params, "actualPrice", rcv.Price)
	params = AppendParameter(params, "imgUrl", rcv.ImgUrl)
	params = AppendParameter(params, "link", rcv.Link)

	return params
}
