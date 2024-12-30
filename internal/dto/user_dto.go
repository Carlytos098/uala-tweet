package dto

type UserDTO struct {
	ID      string   `json:"id"`
	Follows []string `json:"follows"`
}
