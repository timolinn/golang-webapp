package viewmodel

import (
	"fmt"
	"golang-webapp/src/model"
)

// Shop data context
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category Data Context
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVm(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVm(category model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", category.ID),
		ImageURL:    category.ImageURL,
		Title:       category.Title,
		Description: category.Description,
	}
}
