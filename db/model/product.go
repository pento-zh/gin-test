package model

type Product struct {
	Common
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
