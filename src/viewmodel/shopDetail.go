package viewmodel

import (
	"golang-webapp/src/model"
)

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply | Shop Detail",
		Active:   "shop",
		Products: []Product{},
	}

	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}

	return result
}
