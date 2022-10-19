package signUp

type userSupDTO struct {
	Login,
	Password string
}

type userCreateDTO struct {
	Login,
	PasswordHash string
}
type tokensDTO struct {
	AtUuid       string
	AccessToken  string
	AtExpires    int64
	RtUuid       string
	RefreshToken string
	RtExpires    int64
}
type signUpDTO struct {
	Uuid   string
	Tokens tokensDTO
}
