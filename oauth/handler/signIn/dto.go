package signIn

import validation "github.com/go-ozzo/ozzo-validation"

type userInDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (d userInDTO) validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Login, validation.Required, validation.Length(5, 30)),
		validation.Field(&d.Password, validation.Required, validation.Length(8, 30)),
	)
}

type userDTO struct {
	uuid         string
	login        string
	passwordHash string
}

type tokensDTO struct {
	AtUuid       string
	AccessToken  string
	AtExpires    int64
	RtUuid       string
	RefreshToken string
	RtExpires    int64
}
type signInDTO struct {
	uuid   string
	tokens tokensDTO
}
