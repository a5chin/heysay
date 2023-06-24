package entity

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Roles []*Role `json:"roles"`
	Email string  `json:"email"`
}
