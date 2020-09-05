package users

type User struct {
	ID       string `json:"userId"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password,,omitempty"`
}
