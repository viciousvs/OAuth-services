package user

//User
type User struct {
	UUID         string `json:"uuid,omitempty"`
	Login        string `json:"login,omitempty"`
	PasswordHash string `json:"hashPassword,omitempty"`
}

//NewUser
func NewUser(uuid, login, passwordHash string) User {
	return User{
		UUID:         uuid,
		Login:        login,
		PasswordHash: passwordHash,
	}
}
