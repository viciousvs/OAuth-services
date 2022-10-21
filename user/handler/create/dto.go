package create

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viciousvs/OAuth-services/user/utils/uuid"
)

type _DTO struct {
	UUID         string `json:"uuid"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

func newDTOwithUUID(login, passwordHash string) _DTO {
	return _DTO{
		uuid.GenerateUUID(),
		login,
		passwordHash,
	}
}

func (d _DTO) validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.UUID, validation.Required),
		validation.Field(&d.Login, validation.Required, validation.Length(3, 30)),
		validation.Field(&d.PasswordHash, validation.Required, validation.Length(8, 100)),
	)
}
