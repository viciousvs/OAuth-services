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
	AccessToken  string
	AtExpires    int64
	RefreshToken string
	RtExpires    int64
}
type signUpDTO struct {
	UUID   string
	Tokens tokensDTO
}
