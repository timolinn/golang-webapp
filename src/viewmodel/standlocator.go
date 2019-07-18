package viewmodel

type Standlocator struct {
	Title  string
	Active string
}

func NewStandlocator() Standlocator {
	result := Standlocator{
		Title:  "Lemonade Stand Supply | Shop Detail",
		Active: "standlocator",
	}

	return result
}
