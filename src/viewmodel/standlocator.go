package viewmodel

type Standlocator struct {
	Title  string
	Active string
}

func NewStandlocator() Standlocator {
	result := Standlocator{
		Title:  "Lemonade Stand Supply - Stand Locator",
		Active: "standlocator",
	}

	return result
}

type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
