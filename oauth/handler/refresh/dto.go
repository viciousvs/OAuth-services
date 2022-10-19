package refresh

type tokensDTO struct {
	AtUuid       string
	AccessToken  string
	AtExpires    int64
	RtUuid       string
	RefreshToken string
	RtExpires    int64
}
