package token

type Token struct {
	AtUUID      string
	AccessToken string
	AtExpire    int64

	Rt
}
