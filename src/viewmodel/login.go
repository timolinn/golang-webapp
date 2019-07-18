package viewmodel

type Login struct {
	Title    string
	Active   string
	Email    string
	Password string
}

func NewLogin() Login {
	result := Login{
		Title:  "Lemonade Stand Supply",
		Active: "home",
	}
	return result
}
