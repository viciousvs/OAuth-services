package create

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viciousvs/OAuth-services/user/utils/hashPassword"
	"github.com/viciousvs/OAuth-services/user/utils/uuid"
)

//createUserDTO
type createUserDTO struct {
	UUID         string
	Login        string
	PasswordHash string
}

//Validate
func (u createUserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Login, validation.Required, validation.Length(3, 255)),
		validation.Field(&u.PasswordHash, validation.Required, validation.Length(8, 255)),
	)
}

//makeCreateUserDTO
func makeCreateUserDTO(login, password string) (createUserDTO, error) {
	hash, err := hashPassword.Hash(password)
	if err != nil {
		return createUserDTO{}, err
	}
	return createUserDTO{
		UUID:         uuid.GenerateUUID(),
		Login:        login,
		PasswordHash: hash,
	}, nil
}
