package model

type User struct {
	UUID         string `json:"uuid,omitempty"`
	Login        string `json:"login,omitempty"`
	PasswordHash string `json:"passwordHash,omitempty"`
}
