package getByLogin

import validation "github.com/go-ozzo/ozzo-validation"

type _DTO struct {
	Login        string
	PasswordHash string
}

func (d _DTO) validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Login, validation.Required, validation.Length(5, 30)),
		validation.Field(&d.PasswordHash, validation.Required, validation.Length(8, 30)),
	)
}
