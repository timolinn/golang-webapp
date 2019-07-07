package viewmodel

type Base struct {
	Title string
}

type Shop struct {
}

type Category struct {
}

func NewBase() Base {
	return Base{
		Title: "Lemonade Stand Supply",
	}
}
