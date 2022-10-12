package getUUID

import validation "github.com/go-ozzo/ozzo-validation"

//getUserDTO
type getUserDTO struct {
	Login    string
	Password string
}

//makeGetUUIDDTO
func makeGetUserDTO(login, password string) getUserDTO {
	return getUserDTO{
		Login:    login,
		Password: password,
	}
}

func (u getUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Login, validation.Required, validation.Length(3, 255)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 255)),
	)
}

func empty(field string) bool {
	return len(field) == 0
}
