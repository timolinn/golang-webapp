package viewmodel

// Base data context
type Base struct {
	Active string
	Title  string
}

// NewBase is Base Struct constructor function
func NewBase() Base {
	return Base{
		Active: "",
		Title:  "Lemonade Stand Supply",
	}
}
