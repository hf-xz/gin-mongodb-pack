package models

type User struct {
	ID
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	SoftDelete
}
