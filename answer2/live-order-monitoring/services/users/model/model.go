package model

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	RoleID       int    `json:"role_id"`
}
