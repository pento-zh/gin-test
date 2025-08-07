package model

type Order struct {
	Common
	ProductId string `json:"product_id"`
	Amount    int    `json:"amount"`
}
