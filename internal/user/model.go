package user

type Account struct {
	ID string `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Password string `json:"password,omitempty"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Token string `json:"token"`
}