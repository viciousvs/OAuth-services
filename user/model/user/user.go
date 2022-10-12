package user

//User
type User struct {
	UUID         string `json:"uuid,omitempty"`
	Login        string `json:"login,omitempty"`
	PasswordHash string `json:"hashPassword,omitempty"`
}

//MakeUser
func MakeUser(uuid, login, passwordHash string) User {
	return User{
		UUID:         uuid,
		Login:        login,
		PasswordHash: passwordHash,
	}
}

//GettingUserDTO
type GettingUserDTO struct {
	UUID  string
	Login string
	//if we have another fields
}
