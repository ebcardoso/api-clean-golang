package dto

type UserDTO struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	IsBlocked bool   `json:"isBlocked,omitempty"`
}
